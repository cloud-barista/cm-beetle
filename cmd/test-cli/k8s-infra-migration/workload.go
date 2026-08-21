// Workload verification: prove the migrated cluster is actually usable, not merely created.
//
// This mirrors what cmd/test-cli/infra does for VM migration, where SSH connectivity is checked
// on every migrated VM. The K8s equivalent is reaching the cluster's API server and running a
// real workload on it.
//
// Beetle exposes no kubeconfig API, so CB-Tumblebug is called directly — the same arrangement
// the VM CLI uses for its SSH check. Headlamp (used by cb-tumblebug's own k8s-test) is not
// needed: it only forwards the bearer token to the API server unchanged, so calling the API
// server directly does the same work without a Docker dependency.
package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"
)

// kubeconfig captures only the fields needed to reach the API server directly.
type kubeconfig struct {
	Clusters []struct {
		Name    string `yaml:"name"`
		Cluster struct {
			Server                   string `yaml:"server"`
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
		} `yaml:"cluster"`
	} `yaml:"clusters"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			Token                 string      `yaml:"token"`
			ClientCertificateData string      `yaml:"client-certificate-data"`
			ClientKeyData         string      `yaml:"client-key-data"`
			Exec                  interface{} `yaml:"exec"`
		} `yaml:"user"`
	} `yaml:"users"`
}

// authMethod is how a kubeconfig expects the caller to authenticate.
type authMethod int

const (
	authStaticToken authMethod = iota // kubeconfig carries a bearer token
	authClientCert                    // kubeconfig carries a client certificate/key pair
	authExecPlugin                    // kubeconfig delegates to an exec credential plugin
)

func (a authMethod) String() string {
	switch a {
	case authStaticToken:
		return "static token in kubeconfig"
	case authClientCert:
		return "client certificate in kubeconfig"
	default:
		return "exec credential plugin"
	}
}

// nginxDeploymentName is fixed: the workload lives inside one cluster and is removed at the end
// of the step, so it cannot collide with another target's run.
const (
	nginxDeploymentName = "beetle-test-nginx"
	nginxServiceName    = "beetle-test-nginx-lb"
)

// stepWorkload fetches the cluster's kubeconfig, reaches its API server, and runs a real
// workload on it. A cluster that reports Active but cannot schedule a pod has not actually
// been migrated in any useful sense, which is what this step is here to catch.
func stepWorkload(_ *resty.Client, cfg TestConfig, auth AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 5, Name: "Workload verification (kubeconfig -> K8s API -> nginx)", StartTime: time.Now()}

	if auth.TumblebugEndpoint == "" {
		res.Duration = time.Since(res.StartTime)
		res.Skipped, res.Success = true, true
		res.Notes = append(res.Notes, "ℹ️  skipped: tumblebugEndpoint not set in auth config")
		return res
	}

	tb := resty.New().SetTimeout(60 * time.Second).SetLogger(restyNoopLogger{})
	if auth.TumblebugApiUsername != "" {
		tb.SetBasicAuth(auth.TumblebugApiUsername, auth.TumblebugApiPassword)
	}
	base := fmt.Sprintf("%s/tumblebug/ns/%s/k8sCluster/%s", auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, report.ClusterID)

	// 1. Kubeconfig. The dedicated sub-resource is required: the generic cluster GET returns
	// AccessInfo.Kubeconfig without the native flag, which never becomes ready on some CSPs.
	// It also lags the Active status, so it is polled rather than read once.
	kcYAML, err := fetchKubeconfig(tb, base+"/kubeconfig", cfg, &res)
	if err != nil {
		res.Duration = time.Since(res.StartTime)
		res.Error = err.Error()
		return res
	}

	access, err := parseKubeconfig(kcYAML)
	if err != nil {
		res.Duration = time.Since(res.StartTime)
		res.Error = err.Error()
		return res
	}
	server := access.Server
	res.Notes = append(res.Notes,
		fmt.Sprintf("✅ kubeconfig obtained (server: %s)", server),
		fmt.Sprintf("ℹ️  auth method: %s", access.Auth))

	// 2. Credential. Only an exec-plugin kubeconfig needs Tumblebug's token endpoint, and that
	// endpoint is not implemented for every CSP — asking for it unconditionally fails on the
	// ones whose kubeconfig already carries the credential.
	if access.Auth == authExecPlugin {
		token, err := fetchToken(tb, base+"/token")
		if err != nil {
			res.Duration = time.Since(res.StartTime)
			res.Error = err.Error()
			return res
		}
		access.Token = token
		res.Notes = append(res.Notes, "✅ cluster token obtained from Tumblebug")
	}

	// 3. API client. The API server presents a certificate signed by the cluster's own CA, so
	// that CA must be trusted explicitly.
	k8s := resty.New().SetTimeout(60 * time.Second).SetLogger(restyNoopLogger{})
	if access.CAPEM != "" {
		k8s.SetRootCertificateFromString(access.CAPEM)
	}
	if access.Auth == authClientCert {
		cert, certErr := tls.X509KeyPair([]byte(access.CertPEM), []byte(access.KeyPEM))
		if certErr != nil {
			res.Duration = time.Since(res.StartTime)
			res.Error = fmt.Sprintf("failed to load client certificate from kubeconfig: %s", certErr)
			return res
		}
		k8s.SetCertificates(cert)
	} else {
		k8s.SetAuthToken(access.Token)
	}

	// 4. Reachability, and node count cross-checked against the recommendation from inside the
	// cluster — the same claim step 4 makes from Tumblebug's view, verified independently.
	if err := checkAPIServer(k8s, server, report, &res); err != nil {
		res.Duration = time.Since(res.StartTime)
		res.Error = err.Error()
		return res
	}

	// 5. Real workload, and — when enabled — reachability from outside the cluster.
	if err := runNginxWorkload(k8s, server, cfg, &res); err != nil {
		res.Duration = time.Since(res.StartTime)
		res.Error = err.Error()
		// Fall through to cleanup: half-created objects should not outlive the step.
		deleteNginx(k8s, server, &res)
		return res
	}
	if cfg.Workload.LoadBalancerEnabled {
		if err := exposeAndVerifyNginx(k8s, server, cfg, &res); err != nil {
			res.Duration = time.Since(res.StartTime)
			res.Error = err.Error()
			deleteNginx(k8s, server, &res)
			return res
		}
	}
	deleteNginx(k8s, server, &res)

	res.Duration = time.Since(res.StartTime)
	res.Success = true
	return res
}

// fetchKubeconfig polls until the kubeconfig is ready; it becomes available some time after the
// cluster reports Active, because the control-plane endpoint is still being provisioned.
func fetchKubeconfig(tb *resty.Client, url string, cfg TestConfig, res *StepResult) (string, error) {
	deadline := time.Now().Add(time.Duration(cfg.Workload.KubeconfigTimeoutSec) * time.Second)
	interval := time.Duration(cfg.Workload.KubeconfigPollSec) * time.Second

	for attempt := 1; ; attempt++ {
		resp, err := tb.R().Get(url)
		if err == nil && resp.StatusCode() == http.StatusOK {
			var body struct {
				Kubeconfig string `json:"kubeconfig"`
			}
			if jsonErr := json.Unmarshal(resp.Body(), &body); jsonErr == nil &&
				strings.Contains(body.Kubeconfig, "apiVersion") {
				return body.Kubeconfig, nil
			}
		}
		if time.Now().After(deadline) {
			return "", fmt.Errorf("kubeconfig not ready within %ds", cfg.Workload.KubeconfigTimeoutSec)
		}
		progressf(res.Target, "... kubeconfig not ready yet, retrying in %v (attempt %d)", interval, attempt)
		time.Sleep(interval)
	}
}

// clusterAccess is everything needed to talk to a cluster's API server.
type clusterAccess struct {
	Server  string
	CAPEM   string
	Auth    authMethod
	Token   string // set when Auth == authStaticToken
	CertPEM string // set when Auth == authClientCert
	KeyPEM  string // set when Auth == authClientCert
}

// parseKubeconfig extracts the API server endpoint, its CA, and how to authenticate.
//
// Which credential a kubeconfig carries is CSP-specific: EKS/GKE/NKS delegate to an exec
// credential plugin and leave no usable secret in the file, whereas AKS and others embed the
// credential directly. Reading it from the file is what tells the caller whether Tumblebug's
// token endpoint is needed at all.
func parseKubeconfig(kcYAML string) (clusterAccess, error) {
	var access clusterAccess

	var kc kubeconfig
	if err := yaml.Unmarshal([]byte(kcYAML), &kc); err != nil {
		return access, fmt.Errorf("failed to parse kubeconfig: %w", err)
	}
	if len(kc.Clusters) == 0 || kc.Clusters[0].Cluster.Server == "" {
		return access, fmt.Errorf("kubeconfig carries no cluster server URL")
	}
	access.Server = strings.TrimSuffix(kc.Clusters[0].Cluster.Server, "/")

	if data := kc.Clusters[0].Cluster.CertificateAuthorityData; data != "" {
		decoded, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			return access, fmt.Errorf("failed to decode certificate-authority-data: %w", err)
		}
		access.CAPEM = string(decoded)
	}

	if len(kc.Users) == 0 {
		return access, fmt.Errorf("kubeconfig carries no user entry")
	}
	u := kc.Users[0].User

	switch {
	case u.Token != "":
		access.Auth = authStaticToken
		access.Token = u.Token
	case u.ClientCertificateData != "" && u.ClientKeyData != "":
		certPEM, err := base64.StdEncoding.DecodeString(u.ClientCertificateData)
		if err != nil {
			return access, fmt.Errorf("failed to decode client-certificate-data: %w", err)
		}
		keyPEM, err := base64.StdEncoding.DecodeString(u.ClientKeyData)
		if err != nil {
			return access, fmt.Errorf("failed to decode client-key-data: %w", err)
		}
		access.Auth = authClientCert
		access.CertPEM = string(certPEM)
		access.KeyPEM = string(keyPEM)
	case u.Exec != nil:
		access.Auth = authExecPlugin
	default:
		return access, fmt.Errorf("kubeconfig user entry carries no recognised credential")
	}

	return access, nil
}

func fetchToken(tb *resty.Client, url string) (string, error) {
	resp, err := tb.R().Get(url)
	if err != nil {
		return "", fmt.Errorf("token request failed: %w", err)
	}
	if resp.StatusCode() != http.StatusOK {
		return "", fmt.Errorf("token request returned HTTP %d: %s", resp.StatusCode(), truncate(string(resp.Body()), 200))
	}
	var body struct {
		ExecCredential struct {
			Status struct {
				Token string `json:"token"`
			} `json:"status"`
		} `json:"execCredential"`
	}
	if err := json.Unmarshal(resp.Body(), &body); err != nil {
		return "", fmt.Errorf("failed to parse token response: %w", err)
	}
	if body.ExecCredential.Status.Token == "" {
		return "", fmt.Errorf("token response carried no token")
	}
	return body.ExecCredential.Status.Token, nil
}

func checkAPIServer(k8s *resty.Client, server string, report *CSPTestReport, res *StepResult) error {
	verResp, err := k8s.R().Get(server + "/version")
	if err != nil {
		return fmt.Errorf("cannot reach the K8s API server: %w", err)
	}
	if verResp.StatusCode() != http.StatusOK {
		return fmt.Errorf("K8s API /version returned HTTP %d: %s", verResp.StatusCode(), truncate(string(verResp.Body()), 200))
	}
	var version struct {
		GitVersion string `json:"gitVersion"`
	}
	_ = json.Unmarshal(verResp.Body(), &version)
	res.Notes = append(res.Notes, fmt.Sprintf("✅ API server reachable (%s)", version.GitVersion))

	nodesResp, err := k8s.R().Get(server + "/api/v1/nodes")
	if err != nil || nodesResp.StatusCode() != http.StatusOK {
		return fmt.Errorf("cannot list nodes via the K8s API")
	}
	var nodes struct {
		Items []struct {
			Status struct {
				Conditions []struct {
					Type   string `json:"type"`
					Status string `json:"status"`
				} `json:"conditions"`
			} `json:"status"`
		} `json:"items"`
	}
	if err := json.Unmarshal(nodesResp.Body(), &nodes); err != nil {
		return fmt.Errorf("failed to parse node list: %w", err)
	}

	ready := 0
	for _, n := range nodes.Items {
		for _, c := range n.Status.Conditions {
			if c.Type == "Ready" && c.Status == "True" {
				ready++
			}
		}
	}

	want := 0
	if report.Recommendation != nil {
		for _, ng := range report.Recommendation.TargetK8sCluster.K8sNodeGroupList {
			want += ng.DesiredNodeSize
		}
	}
	if want > 0 && ready != want {
		res.Notes = append(res.Notes, fmt.Sprintf("❌ %d/%d node(s) Ready, recommendation asked for %d", ready, len(nodes.Items), want))
		return fmt.Errorf("cluster has %d ready node(s), expected %d", ready, want)
	}
	res.Notes = append(res.Notes, fmt.Sprintf("✅ %d node(s) Ready, matching the recommendation", ready))
	return nil
}

// runNginxWorkload creates a Deployment and waits for its pod to run. Scope stops at Running:
// a LoadBalancer Service would add per-CSP provisioning time as another variable.
func runNginxWorkload(k8s *resty.Client, server string, cfg TestConfig, res *StepResult) error {
	deployment := map[string]interface{}{
		"apiVersion": "apps/v1",
		"kind":       "Deployment",
		"metadata":   map[string]interface{}{"name": nginxDeploymentName},
		"spec": map[string]interface{}{
			"replicas": 1,
			"selector": map[string]interface{}{"matchLabels": map[string]string{"app": nginxDeploymentName}},
			"template": map[string]interface{}{
				"metadata": map[string]interface{}{"labels": map[string]string{"app": nginxDeploymentName}},
				"spec": map[string]interface{}{
					"containers": []map[string]interface{}{{
						"name":  "nginx",
						"image": "nginx:stable-alpine",
						"ports": []map[string]interface{}{{"containerPort": 80}},
					}},
				},
			},
		},
	}

	resp, err := k8s.R().
		SetHeader("Content-Type", "application/json").
		SetBody(deployment).
		Post(server + "/apis/apps/v1/namespaces/default/deployments")
	if err != nil {
		return fmt.Errorf("failed to create nginx Deployment: %w", err)
	}
	if resp.StatusCode() != http.StatusCreated && resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("nginx Deployment creation returned HTTP %d: %s", resp.StatusCode(), truncate(string(resp.Body()), 200))
	}
	res.Notes = append(res.Notes, "✅ nginx Deployment created")

	deadline := time.Now().Add(time.Duration(cfg.Workload.PodReadyTimeoutSec) * time.Second)
	interval := time.Duration(cfg.Workload.PodPollSec) * time.Second
	podURL := fmt.Sprintf("%s/api/v1/namespaces/default/pods?labelSelector=app%%3D%s", server, nginxDeploymentName)

	for attempt := 1; ; attempt++ {
		time.Sleep(interval)

		podsResp, podErr := k8s.R().Get(podURL)
		if podErr == nil && podsResp.StatusCode() == http.StatusOK {
			var pods struct {
				Items []struct {
					Status struct {
						Phase string `json:"phase"`
					} `json:"status"`
				} `json:"items"`
			}
			if json.Unmarshal(podsResp.Body(), &pods) == nil {
				for _, p := range pods.Items {
					if p.Status.Phase == "Running" {
						res.Notes = append(res.Notes, fmt.Sprintf("✅ nginx pod Running (attempt %d)", attempt))
						return nil
					}
				}
				if len(pods.Items) > 0 {
					progressf(res.Target, "... nginx pod phase: %s (attempt %d)", pods.Items[0].Status.Phase, attempt)
				}
			}
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("nginx pod did not reach Running within %ds", cfg.Workload.PodReadyTimeoutSec)
		}
	}
}

// exposeAndVerifyNginx publishes the workload through a LoadBalancer Service and fetches it
// from outside the cluster. Scheduling a pod proves the cluster can run work; this proves the
// work is actually reachable, which is a distinct failure mode — a cluster can come up healthy
// and still fail to provision a working load balancer.
func exposeAndVerifyNginx(k8s *resty.Client, server string, cfg TestConfig, res *StepResult) error {
	service := map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "Service",
		"metadata":   map[string]interface{}{"name": nginxServiceName},
		"spec": map[string]interface{}{
			"type":     "LoadBalancer",
			"selector": map[string]string{"app": nginxDeploymentName},
			"ports": []map[string]interface{}{{
				"port": 80, "targetPort": 80, "protocol": "TCP",
			}},
		},
	}

	resp, err := k8s.R().
		SetHeader("Content-Type", "application/json").
		SetBody(service).
		Post(server + "/api/v1/namespaces/default/services")
	if err != nil {
		return fmt.Errorf("failed to create LoadBalancer Service: %w", err)
	}
	if resp.StatusCode() != http.StatusCreated && resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("LoadBalancer Service creation returned HTTP %d: %s",
			resp.StatusCode(), truncate(string(resp.Body()), 200))
	}
	res.Notes = append(res.Notes, "✅ LoadBalancer Service created")

	address, err := waitForLoadBalancerAddress(k8s, server, cfg, res)
	if err != nil {
		return err
	}
	res.Notes = append(res.Notes, fmt.Sprintf("✅ LoadBalancer address assigned: %s", address))

	return fetchThroughLoadBalancer(address, cfg, res)
}

// waitForLoadBalancerAddress polls the Service until the CSP assigns an ingress address. AWS
// publishes a hostname, other CSPs an IP, so both fields are accepted.
func waitForLoadBalancerAddress(k8s *resty.Client, server string, cfg TestConfig, res *StepResult) (string, error) {
	deadline := time.Now().Add(time.Duration(cfg.Workload.LbAddressTimeoutSec) * time.Second)
	interval := time.Duration(cfg.Workload.LbPollSec) * time.Second
	url := server + "/api/v1/namespaces/default/services/" + nginxServiceName

	for attempt := 1; ; attempt++ {
		time.Sleep(interval)

		resp, err := k8s.R().Get(url)
		if err == nil && resp.StatusCode() == http.StatusOK {
			var svc struct {
				Status struct {
					LoadBalancer struct {
						Ingress []struct {
							IP       string `json:"ip"`
							Hostname string `json:"hostname"`
						} `json:"ingress"`
					} `json:"loadBalancer"`
				} `json:"status"`
			}
			if json.Unmarshal(resp.Body(), &svc) == nil {
				for _, ing := range svc.Status.LoadBalancer.Ingress {
					if ing.IP != "" {
						return ing.IP, nil
					}
					if ing.Hostname != "" {
						return ing.Hostname, nil
					}
				}
			}
		}
		if time.Now().After(deadline) {
			return "", fmt.Errorf("LoadBalancer address was not assigned within %ds", cfg.Workload.LbAddressTimeoutSec)
		}
		progressf(res.Target, "... waiting for LoadBalancer address (attempt %d)", attempt)
	}
}

// fetchThroughLoadBalancer retries until the endpoint serves. An address appears before the
// load balancer is actually forwarding traffic — health checks must pass first, and an AWS
// hostname additionally has to resolve — so the first attempts are expected to fail.
func fetchThroughLoadBalancer(address string, cfg TestConfig, res *StepResult) error {
	client := resty.New().SetTimeout(15 * time.Second).SetLogger(restyNoopLogger{})
	url := "http://" + address + "/"

	deadline := time.Now().Add(time.Duration(cfg.Workload.LbAccessTimeoutSec) * time.Second)
	interval := time.Duration(cfg.Workload.LbPollSec) * time.Second
	lastErr := "no attempt made"

	for attempt := 1; ; attempt++ {
		resp, err := client.R().Get(url)
		switch {
		case err != nil:
			lastErr = err.Error()
		case resp.StatusCode() == http.StatusOK:
			res.Notes = append(res.Notes,
				fmt.Sprintf("✅ nginx served over the LoadBalancer at %s (attempt %d)", url, attempt))
			return nil
		default:
			lastErr = fmt.Sprintf("HTTP %d", resp.StatusCode())
		}

		if time.Now().After(deadline) {
			return fmt.Errorf("LoadBalancer at %s did not serve within %ds (last: %s)",
				url, cfg.Workload.LbAccessTimeoutSec, lastErr)
		}
		progressf(res.Target, "... LoadBalancer not serving yet: %s (attempt %d)", lastErr, attempt)
		time.Sleep(interval)
	}
}

// deleteNginx removes the Service first, then the Deployment.
//
// Order matters: a LoadBalancer Service holds a real, billable cloud load balancer, and deleting
// it is what releases that. A failure here is reported prominently rather than shrugged off —
// unlike the Deployment, a leaked load balancer keeps costing money after the run.
func deleteNginx(k8s *resty.Client, server string, res *StepResult) {
	svcResp, svcErr := k8s.R().Delete(server + "/api/v1/namespaces/default/services/" + nginxServiceName)
	switch {
	case svcErr != nil:
		res.Notes = append(res.Notes, fmt.Sprintf("❌ LoadBalancer Service cleanup failed (%s) — check the CSP for a leaked load balancer", svcErr))
	case svcResp.StatusCode() == http.StatusNotFound:
		// Never created (LoadBalancer verification disabled or the Service creation failed).
	case svcResp.StatusCode() != http.StatusOK:
		res.Notes = append(res.Notes, fmt.Sprintf("❌ LoadBalancer Service cleanup returned HTTP %d — check the CSP for a leaked load balancer", svcResp.StatusCode()))
	default:
		res.Notes = append(res.Notes, "✅ LoadBalancer Service removed")
	}

	resp, err := k8s.R().Delete(server + "/apis/apps/v1/namespaces/default/deployments/" + nginxDeploymentName)
	if err != nil || (resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusNotFound) {
		res.Notes = append(res.Notes, "ℹ️  nginx Deployment cleanup did not confirm; the cluster is deleted next anyway")
		return
	}
	res.Notes = append(res.Notes, "✅ nginx Deployment removed")
}

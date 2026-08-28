// Package main is the starting point of CM-Beetle Managed RDBMS Test CLI
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
)

// restyNoopLogger silences all Resty log output.
type restyNoopLogger struct{}

func (restyNoopLogger) Errorf(_ string, _ ...interface{}) {}
func (restyNoopLogger) Warnf(_ string, _ ...interface{})  {}
func (restyNoopLogger) Debugf(_ string, _ ...interface{}) {}

// ============================================================================
// Configuration & Type Definitions
// ============================================================================

// TestConfig holds test configuration loaded from YAML.
type TestConfig struct {
	Test struct {
		Set struct {
			Mode string `yaml:"mode"` // sequential or parallel
		} `yaml:"set"`
		Cases []TestCase `yaml:"cases"`
	} `yaml:"test"`
	Beetle struct {
		Endpoint        string `yaml:"endpoint"`
		NamespaceID     string `yaml:"namespaceId"`
		RequestBodyFile string `yaml:"requestBodyFile"`
		AuthConfigFile  string `yaml:"authConfigFile"`
	} `yaml:"beetle"`
}

// SubnetConfig represents an individual subnet within a VNet.
type SubnetConfig struct {
	Name string `yaml:"name" json:"name"`
	CIDR string `yaml:"cidr" json:"ipv4_CIDR"`
	Zone string `yaml:"zone" json:"zone"`
}

// TestCase holds parameters for testing a single CSP.
type TestCase struct {
	Name                     string         `yaml:"name"`
	Csp                      string         `yaml:"csp"`
	Region                   string         `yaml:"region"`
	RdbmsId                  string         `yaml:"rdbmsId"`
	ConnectionName           string         `yaml:"connectionName"`
	VNetName                 string         `yaml:"vNetName"`
	CidrBlock                string         `yaml:"cidrBlock"`
	Subnets                  []SubnetConfig `yaml:"subnets"`
	SecurityGroupName        string         `yaml:"securityGroupName"`
	AutoFillDefaults         bool           `yaml:"autoFillDefaults"`
	CommonSpec               string         `yaml:"commonSpec"`
	CommonImage              string         `yaml:"commonImage"`
	VmSpecId                 string         `yaml:"vmSpecId"`
	VmImageId                string         `yaml:"vmImageId"`
	VmOSType                 string         `yaml:"vmOSType"`
	VmvCPU                   string         `yaml:"vmvCPU"`
	VmMemoryGiB              string         `yaml:"vmMemoryGiB"`
	DBEngine                 string         `yaml:"dbEngine"`
	DBEngineVersion          string         `yaml:"dbEngineVersion"`
	DBInstanceSpec           string         `yaml:"dbInstanceSpec"`
	StorageType              string         `yaml:"storageType"`
	StorageSize              int            `yaml:"storageSize"`
	AdminUserName            string         `yaml:"adminUserName"`
	AdminUserPassword        string         `yaml:"adminUserPassword"`
	PublicAccess             bool           `yaml:"publicAccess"`
	HighAvailability         bool           `yaml:"highAvailability"`
	DatabaseName             string         `yaml:"databaseName"`
	NHNDBSGToAllowAllInbound bool           `yaml:"nhnDBSGToAllowAllInbound"`
	ExternalDataIOTest       bool           `yaml:"externalDataIOTest"`
	InternalDataIOTest       bool           `yaml:"internalDataIOTest"`
	Execute                  bool           `yaml:"execute"`
}

// AuthConfig holds basic auth credentials.
type AuthConfig struct {
	BeetleApiUsername    string `json:"beetleApiUsername"`
	BeetleApiPassword    string `json:"beetleApiPassword"`
	TumblebugApiUsername string `json:"tumblebugApiUsername"`
	TumblebugApiPassword string `json:"tumblebugApiPassword"`
	TumblebugEndpoint    string `json:"tumblebugEndpoint"`
}

// RDBMSRequestFile represents the source RDBMS JSON file.
type RDBMSRequestFile struct {
	NameSeed             string                            `json:"nameSeed"`
	SourceRDBMSInstances []rdbmsmodel.SourceRDBMSProperty `json:"sourceRDBMSInstances"`
}

// TestResults holds the result of a single step.
type TestResults struct {
	TestName     string        `json:"testName"`
	StartTime    time.Time     `json:"startTime"`
	EndTime      time.Time     `json:"endTime"`
	Duration     time.Duration `json:"duration"`
	Success      bool          `json:"success"`
	Skipped      bool          `json:"skipped"`
	StatusCode   int           `json:"statusCode"`
	Response     interface{}   `json:"response,omitempty"`
	Error        string        `json:"error,omitempty"`
	ErrorMessage string        `json:"errorMessage,omitempty"`
	RequestURL   string        `json:"requestUrl,omitempty"`
	RequestBody  interface{}   `json:"requestBody,omitempty"`
}

// RDBMSTestReport holds all results and traces for a single CSP test.
type RDBMSTestReport struct {
	CSP                    string
	Region                 string
	DisplayName            string
	TestDate               string
	TestTime               string
	TestDateTime           time.Time
	BeetleURL              string
	TumblebugURL           string
	NamespaceID            string
	NameSeed               string
	SourceRDBMS            RDBMSRequestFile
	RecommendationResponse interface{}
	ValidationResponse     interface{}
	MigrationResponse      interface{}
	ListResponse           interface{}
	GetResponse            interface{}
	DatabaseListResponse   interface{}
	ExternalDataIOTest     string
	InternalDataIOTest     string
	TestResults            []TestResults
	Summary                TestResults
}

// ============================================================================
// Main Execution
// ============================================================================

func main() {
	configPath := flag.String("config", "cmd/test-cli/rdbms/testconf/test-config.yaml", "Path to test configuration YAML")
	parallelFlag := flag.Bool("parallel", false, "Run test cases in parallel")
	dryRunFlag := flag.Bool("dry-run", false, "Parse config and validate without executing tests")
	nsIdOverride := flag.String("nsId", "", "Override namespace ID")
	flag.Parse()

	logger.NewLogger(logger.Config{LogLevel: "info"})
	log.Info().Msg("Starting CM-Beetle Managed RDBMS Test CLI")

	// 1. Load configuration
	cfgData, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to read config file: %s", *configPath)
	}

	var testConfig TestConfig
	if err := yaml.Unmarshal(cfgData, &testConfig); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse test config YAML")
	}

	if *nsIdOverride != "" {
		testConfig.Beetle.NamespaceID = *nsIdOverride
	}
	if testConfig.Beetle.NamespaceID == "" {
		testConfig.Beetle.NamespaceID = "mig01"
	}
	if *parallelFlag {
		testConfig.Test.Set.Mode = "parallel"
	}

	// 2. Load Auth & Request files
	resolvePath := func(p string) string {
		if _, err := os.Stat(p); err == nil {
			return p
		}
		// Try resolving relative to config file directory
		cfgDir := filepath.Dir(*configPath)
		candidate := filepath.Join(cfgDir, p)
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
		candidate2 := filepath.Join(cfgDir, filepath.Base(p))
		if _, err := os.Stat(candidate2); err == nil {
			return candidate2
		}
		return p
	}

	authPath := resolvePath(testConfig.Beetle.AuthConfigFile)
	authData, err := os.ReadFile(authPath)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to read auth config file: %s (resolved from %s)", authPath, testConfig.Beetle.AuthConfigFile)
	}
	var authConfig AuthConfig
	if err := json.Unmarshal(authData, &authConfig); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse auth config JSON")
	}

	reqPath := resolvePath(testConfig.Beetle.RequestBodyFile)
	reqData, err := os.ReadFile(reqPath)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to read request body file: %s (resolved from %s)", reqPath, testConfig.Beetle.RequestBodyFile)
	}
	var rdbmsReq RDBMSRequestFile
	if err := json.Unmarshal(reqData, &rdbmsReq); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse request body JSON")
	}

	// 3. Filter active test cases (execute: true only)
	activeCases := make([]TestCase, 0)
	for _, tc := range testConfig.Test.Cases {
		if tc.Execute {
			activeCases = append(activeCases, tc)
		}
	}

	log.Info().
		Str("mode", testConfig.Test.Set.Mode).
		Str("namespaceId", testConfig.Beetle.NamespaceID).
		Str("beetleEndpoint", testConfig.Beetle.Endpoint).
		Str("tumblebugEndpoint", authConfig.TumblebugEndpoint).
		Int("totalTestCases", len(testConfig.Test.Cases)).
		Int("activeTestCases", len(activeCases)).
		Msg("Configuration loaded successfully")

	if *dryRunFlag {
		log.Info().Msg("[Dry Run] Configuration parsed and validated successfully. Exiting without execution.")
		return
	}

	if len(activeCases) == 0 {
		log.Warn().Msg("No active test cases found (execute: true). Exiting.")
		return
	}

	// 4. Run test cases (Sequential or Parallel)
	reports := make([]*RDBMSTestReport, len(activeCases))
	startTime := time.Now()

	if strings.ToLower(testConfig.Test.Set.Mode) == "parallel" {
		var wg sync.WaitGroup
		for i, tc := range activeCases {
			wg.Add(1)
			go func(idx int, testCase TestCase) {
				defer wg.Done()
				reports[idx] = runSingleRDBMSTest(testConfig, authConfig, rdbmsReq, testCase)
			}(i, tc)
		}
		wg.Wait()
	} else {
		for i, tc := range activeCases {
			reports[i] = runSingleRDBMSTest(testConfig, authConfig, rdbmsReq, tc)
		}
	}

	totalDuration := time.Since(startTime)

	// 5. Generate Markdown Reports
	outputDir := "testresult"
	_ = os.MkdirAll(outputDir, 0755)

	generateSummaryReport(outputDir, reports, totalDuration)
	for _, r := range reports {
		if r != nil {
			generateDetailedReport(outputDir, r)
		}
	}

	log.Info().
		Str("totalDuration", totalDuration.String()).
		Str("reportDir", outputDir).
		Msg("CM-Beetle Managed RDBMS Test Run Completed")
}

// ============================================================================
// Single Test Case Runner
// ============================================================================

func runSingleRDBMSTest(
	cfg TestConfig,
	auth AuthConfig,
	baseReq RDBMSRequestFile,
	tc TestCase,
) *RDBMSTestReport {
	now := time.Now()
	report := &RDBMSTestReport{
		CSP:          tc.Csp,
		Region:       tc.Region,
		DisplayName:  tc.Name,
		TestDate:     now.Format("2006-01-02"),
		TestTime:     now.Format("15:04:05"),
		TestDateTime: now,
		BeetleURL:    cfg.Beetle.Endpoint,
		TumblebugURL: auth.TumblebugEndpoint,
		NamespaceID:  cfg.Beetle.NamespaceID,
		NameSeed:     baseReq.NameSeed,
		SourceRDBMS:  baseReq,
		TestResults:  make([]TestResults, 0),
	}

	log.Info().Str("csp", tc.Csp).Str("region", tc.Region).Msgf("[%s] ====== STARTING RDBMS TEST ======", tc.Name)

	bClient := createRestClient(auth.BeetleApiUsername, auth.BeetleApiPassword)
	tbClient := createRestClient(auth.TumblebugApiUsername, auth.TumblebugApiPassword)

	var vNetId string
	var subnetIds []string
	var sgIds []string
	var candidateRDBMSId string
	var createdRDBMSId string
	var rdbmsEndpoint string
	dbName := tc.DatabaseName
	if dbName == "" {
		dbName = "sampledb"
	}

	// ------------------------------------------------------------------------
	// Phase 0: Pre-flight Spec & Image Review and Resolution (/specImagePairReview)
	// ------------------------------------------------------------------------
	reviewStep := TestResults{TestName: "Tumblebug POST /specImagePairReview (Pre-flight Spec & Image Review)", StartTime: time.Now()}
	resolveAndReviewSpecAndImage(auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, &tc, tbClient, &reviewStep)
	reviewStep.EndTime = time.Now()
	reviewStep.Duration = reviewStep.EndTime.Sub(reviewStep.StartTime)
	report.TestResults = append(report.TestResults, reviewStep)

	// ------------------------------------------------------------------------
	// Phase 1: Pre-requisite Cloud Infrastructure Provisioning (VNet & SecurityGroup)
	// ------------------------------------------------------------------------
	// 1.1 Create VNet with Subnets via POST /ns/{nsId}/resources/vNet
	vNetStep := TestResults{TestName: "Tumblebug POST /resources/vNet (Create VNet & Subnets)", StartTime: time.Now()}
	vNetName := tc.VNetName
	if vNetName == "" {
		vNetName = fmt.Sprintf("test-rdbms-vnet-%s", tc.Csp)
	}
	cidrBlock := tc.CidrBlock
	if cidrBlock == "" {
		cidrBlock = "10.0.0.0/16"
	}

	subnetReqs := make([]map[string]any, 0, len(tc.Subnets))
	if len(tc.Subnets) > 0 {
		for _, s := range tc.Subnets {
			subnetReqs = append(subnetReqs, map[string]any{
				"name":      s.Name,
				"ipv4_CIDR": s.CIDR,
				"zone":      s.Zone,
			})
		}
	} else {
		// Default 2 subnets across zones for standard SubnetGroup requirements
		subnetReqs = append(subnetReqs,
			map[string]any{"name": "subnet-1", "ipv4_CIDR": "10.0.1.0/24", "zone": ""},
			map[string]any{"name": "subnet-2", "ipv4_CIDR": "10.0.2.0/24", "zone": ""},
		)
	}

	vNetReqBody := map[string]any{
		"name":           vNetName,
		"connectionName": tc.ConnectionName,
		"cidrBlock":      cidrBlock,
		"subnetInfoList": subnetReqs,
		"description":    "Pre-requisite VNet for CM-Beetle RDBMS test",
	}
	vNetURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/vNet", auth.TumblebugEndpoint, cfg.Beetle.NamespaceID)
	log.Info().Msgf("[%s] Creating pre-requisite VNet '%s' (connection=%s)...", tc.Csp, vNetName, tc.ConnectionName)

	vNetResp, vNetErr := tbClient.R().SetBody(vNetReqBody).Post(vNetURL)
	vNetStep.EndTime = time.Now()
	vNetStep.Duration = vNetStep.EndTime.Sub(vNetStep.StartTime)
	vNetStep.RequestURL = vNetURL
	vNetStep.RequestBody = maskSecrets(vNetReqBody)

	if vNetErr != nil || vNetResp.IsError() {
		// If already exists, fetch existing VNet details
		if vNetResp != nil && (vNetResp.StatusCode() == 409 || strings.Contains(vNetResp.String(), "already exists")) {
			log.Info().Msgf("[%s] VNet '%s' already exists; fetching details...", tc.Csp, vNetName)
			getVNetURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/vNet/%s", auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, vNetName)
			getResp, getErr := tbClient.R().Get(getVNetURL)
			if getErr == nil && !getResp.IsError() {
				vNetStep.Success = true
				vNetStep.StatusCode = getResp.StatusCode()
				var vNetInfo map[string]any
				_ = json.Unmarshal(getResp.Body(), &vNetInfo)
				vNetStep.Response = vNetInfo
				if id, ok := vNetInfo["id"].(string); ok {
					vNetId = id
				}
				if sList, ok := vNetInfo["subnetInfoList"].([]any); ok {
					for _, sItem := range sList {
						if sMap, ok := sItem.(map[string]any); ok {
							if sId, ok := sMap["id"].(string); ok {
								subnetIds = append(subnetIds, sId)
							}
						}
					}
				}
				log.Info().Msgf("[%s] Reused existing VNet OK: id=%s, subnets=%v", tc.Csp, vNetId, subnetIds)
			} else {
				vNetStep.Success = false
				vNetStep.StatusCode = vNetResp.StatusCode()
				vNetStep.Error = fmt.Sprintf("err: %v, body: %s", vNetErr, vNetResp.String())
				log.Error().Msgf("[%s] TB Create VNet failed: %s", tc.Csp, vNetStep.Error)
			}
		} else {
			vNetStep.Success = false
			vNetStep.StatusCode = vNetResp.StatusCode()
			vNetStep.Error = fmt.Sprintf("err: %v, body: %s", vNetErr, vNetResp.String())
			log.Error().Msgf("[%s] TB Create VNet failed: %s", tc.Csp, vNetStep.Error)
		}
	} else {
		vNetStep.Success = true
		vNetStep.StatusCode = vNetResp.StatusCode()
		var vNetInfo map[string]any
		_ = json.Unmarshal(vNetResp.Body(), &vNetInfo)
		vNetStep.Response = vNetInfo

		if id, ok := vNetInfo["id"].(string); ok {
			vNetId = id
		}
		if sList, ok := vNetInfo["subnetInfoList"].([]any); ok {
			for _, sItem := range sList {
				if sMap, ok := sItem.(map[string]any); ok {
					if sId, ok := sMap["id"].(string); ok {
						subnetIds = append(subnetIds, sId)
					}
				}
			}
		}
		log.Info().Msgf("[%s] TB Create VNet OK: id=%s, subnets=%v", tc.Csp, vNetId, subnetIds)
	}
	report.TestResults = append(report.TestResults, vNetStep)

	// 1.2 Create SecurityGroup via POST /ns/{nsId}/resources/securityGroup
	sgStep := TestResults{TestName: "Tumblebug POST /resources/securityGroup (Create SecurityGroup)", StartTime: time.Now()}
	if vNetId != "" {
		sgName := tc.SecurityGroupName
		if sgName == "" {
			sgName = fmt.Sprintf("test-rdbms-sg-%s", tc.Csp)
		}
		dbPort := "3306"
		if strings.ToLower(tc.DBEngine) == "postgresql" || strings.ToLower(tc.DBEngine) == "postgres" {
			dbPort = "5432"
		}
		firewallRules := []map[string]any{
			{"Ports": dbPort, "Protocol": "TCP", "Direction": "inbound", "CIDR": "0.0.0.0/0"},
			{"Ports": "22", "Protocol": "TCP", "Direction": "inbound", "CIDR": "0.0.0.0/0"},
		}

		sgReqBody := map[string]any{
			"name":           sgName,
			"connectionName": tc.ConnectionName,
			"vNetId":         vNetId,
			"description":    "Pre-requisite SecurityGroup for CM-Beetle RDBMS test",
			"firewallRules":  firewallRules,
		}
		sgURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/securityGroup", auth.TumblebugEndpoint, cfg.Beetle.NamespaceID)
		log.Info().Msgf("[%s] Creating pre-requisite SecurityGroup '%s' (vNetId=%s)...", tc.Csp, sgName, vNetId)

		sgResp, sgErr := tbClient.R().SetBody(sgReqBody).Post(sgURL)
		sgStep.EndTime = time.Now()
		sgStep.Duration = sgStep.EndTime.Sub(sgStep.StartTime)
		sgStep.RequestURL = sgURL
		sgStep.RequestBody = maskSecrets(sgReqBody)

		if sgErr != nil || sgResp.IsError() {
			if sgResp != nil && (sgResp.StatusCode() == 409 || strings.Contains(sgResp.String(), "already exists")) {
				log.Info().Msgf("[%s] SecurityGroup '%s' already exists; fetching details...", tc.Csp, sgName)
				getSGURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/securityGroup/%s", auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, sgName)
				getResp, getErr := tbClient.R().Get(getSGURL)
				if getErr == nil && !getResp.IsError() {
					sgStep.Success = true
					sgStep.StatusCode = getResp.StatusCode()
					var sgInfo map[string]any
					_ = json.Unmarshal(getResp.Body(), &sgInfo)
					sgStep.Response = sgInfo
					if id, ok := sgInfo["id"].(string); ok {
						sgIds = append(sgIds, id)
					}
					log.Info().Msgf("[%s] Reused existing SecurityGroup OK: id=%v", tc.Csp, sgIds)
				} else {
					sgStep.Success = false
					sgStep.StatusCode = sgResp.StatusCode()
					sgStep.Error = fmt.Sprintf("err: %v, body: %s", sgErr, sgResp.String())
					log.Error().Msgf("[%s] TB Create SecurityGroup failed: %s", tc.Csp, sgStep.Error)
				}
			} else {
				sgStep.Success = false
				sgStep.StatusCode = sgResp.StatusCode()
				sgStep.Error = fmt.Sprintf("err: %v, body: %s", sgErr, sgResp.String())
				log.Error().Msgf("[%s] TB Create SecurityGroup failed: %s", tc.Csp, sgStep.Error)
			}
		} else {
			sgStep.Success = true
			sgStep.StatusCode = sgResp.StatusCode()
			var sgInfo map[string]any
			_ = json.Unmarshal(sgResp.Body(), &sgInfo)
			sgStep.Response = sgInfo

			if id, ok := sgInfo["id"].(string); ok {
				sgIds = append(sgIds, id)
			}
			log.Info().Msgf("[%s] TB Create SecurityGroup OK: id=%v", tc.Csp, sgIds)
		}
	} else {
		sgStep.Skipped = true
		sgStep.ErrorMessage = "Skipped because VNet creation failed"
	}
	report.TestResults = append(report.TestResults, sgStep)

	// ------------------------------------------------------------------------
	// Phase 2: CM-Beetle Recommendation & Validation APIs
	// ------------------------------------------------------------------------
	// 2.1 GET /recommendation/middleware/rdbms/support
	suppStep := TestResults{TestName: "Beetle GET RDBMS Support", StartTime: time.Now()}
	suppURL := fmt.Sprintf("%s/recommendation/middleware/rdbms/support?cspType=%s", cfg.Beetle.Endpoint, tc.Csp)
	suppResp, suppErr := bClient.R().Get(suppURL)
	suppStep.EndTime = time.Now()
	suppStep.Duration = suppStep.EndTime.Sub(suppStep.StartTime)
	suppStep.RequestURL = suppURL
	if suppErr != nil || suppResp.IsError() {
		suppStep.Success = false
		suppStep.StatusCode = suppResp.StatusCode()
		suppStep.Error = fmt.Sprintf("err: %v, body: %s", suppErr, suppResp.String())
	} else {
		suppStep.Success = true
		suppStep.StatusCode = suppResp.StatusCode()
		var suppMap map[string]any
		_ = json.Unmarshal(suppResp.Body(), &suppMap)
		suppStep.Response = suppMap
	}
	report.TestResults = append(report.TestResults, suppStep)

	// 2.2 GET /recommendation/middleware/rdbms/capability
	capStep := TestResults{TestName: "Beetle GET RDBMS Capability", StartTime: time.Now()}
	capURL := fmt.Sprintf("%s/recommendation/middleware/rdbms/capability?connectionName=%s", cfg.Beetle.Endpoint, tc.ConnectionName)
	capResp, capErr := bClient.R().Get(capURL)
	capStep.EndTime = time.Now()
	capStep.Duration = capStep.EndTime.Sub(capStep.StartTime)
	capStep.RequestURL = capURL
	if capErr != nil || capResp.IsError() {
		capStep.Success = false
		capStep.StatusCode = capResp.StatusCode()
		capStep.Error = fmt.Sprintf("err: %v, body: %s", capErr, capResp.String())
	} else {
		capStep.Success = true
		capStep.StatusCode = capResp.StatusCode()
		var capMap map[string]any
		_ = json.Unmarshal(capResp.Body(), &capMap)
		capStep.Response = capMap
	}
	report.TestResults = append(report.TestResults, capStep)

	// 2.3 POST /recommendation/middleware/rdbms
	recStep := TestResults{TestName: "Beetle POST Recommend RDBMS", StartTime: time.Now()}
	recReqBody := controller.RecommendRDBMSRequest{
		DesiredCloud: rdbmsmodel.CloudProperty{
			Csp:    tc.Csp,
			Region: tc.Region,
		},
		SourceRDBMSInstances: baseReq.SourceRDBMSInstances,
	}
	recURL := fmt.Sprintf("%s/recommendation/middleware/rdbms", cfg.Beetle.Endpoint)
	recResp, recErr := bClient.R().SetBody(recReqBody).Post(recURL)
	recStep.EndTime = time.Now()
	recStep.Duration = recStep.EndTime.Sub(recStep.StartTime)
	recStep.RequestURL = recURL
	recStep.RequestBody = maskSecrets(recReqBody)

	var recommendedResult rdbmsmodel.RecommendedRDBMS
	if recErr != nil || recResp.IsError() {
		recStep.Success = false
		recStep.StatusCode = recResp.StatusCode()
		recStep.Error = fmt.Sprintf("err: %v, body: %s", recErr, recResp.String())
		log.Error().Msgf("[%s] Beetle Recommend RDBMS failed: %s", tc.Csp, recStep.Error)
	} else {
		recStep.Success = true
		recStep.StatusCode = recResp.StatusCode()
		var apiResp model.ApiResponse[rdbmsmodel.RecommendedRDBMS]
		if err := json.Unmarshal(recResp.Body(), &apiResp); err == nil {
			recommendedResult = apiResp.Data
			report.RecommendationResponse = recommendedResult
			recStep.Response = recommendedResult
		}
		log.Info().Msgf("[%s] Beetle Recommend RDBMS OK: %d instance(s) recommended", tc.Csp, len(recommendedResult.TargetRDBMSInstances))
	}
	report.TestResults = append(report.TestResults, recStep)

	// Bind created infrastructure IDs & test parameters into recommendation
	if len(recommendedResult.TargetRDBMSInstances) > 0 {
		for i := range recommendedResult.TargetRDBMSInstances {
			inst := &recommendedResult.TargetRDBMSInstances[i]
			inst.RDBMSName = tc.RdbmsId
			inst.VNetId = vNetId
			inst.SubnetIds = subnetIds
			inst.SecurityGroupIds = sgIds
			if tc.AdminUserName != "" {
				inst.AdminUserName = tc.AdminUserName
			}
			if tc.AdminUserPassword != "" {
				inst.AdminUserPassword = tc.AdminUserPassword
			}
			if tc.DBInstanceSpec != "" {
				inst.DBInstanceSpec = tc.DBInstanceSpec
			}
			if tc.StorageType != "" {
				inst.StorageType = tc.StorageType
			}
			if tc.StorageSize > 0 {
				inst.StorageSize = tc.StorageSize
			}
			inst.PublicAccess = tc.PublicAccess
			inst.HighAvailability = tc.HighAvailability
			inst.Databases = []rdbmsmodel.TargetDatabase{
				{DatabaseName: dbName},
			}
		}
	}

	// 2.4 POST /recommendation/middleware/rdbms/validate
	valStep := TestResults{TestName: "Beetle POST Validate RDBMS Recommendation", StartTime: time.Now()}
	if len(recommendedResult.TargetRDBMSInstances) > 0 {
		valReq := rdbmsmodel.RDBMSCreateRequest{
			Name:              recommendedResult.TargetRDBMSInstances[0].RDBMSName,
			ConnectionName:    tc.ConnectionName,
			VNetId:            vNetId,
			SubnetIds:         subnetIds,
			SecurityGroupIds:  sgIds,
			DBEngine:          recommendedResult.TargetRDBMSInstances[0].DBEngine,
			DBEngineVersion:   recommendedResult.TargetRDBMSInstances[0].DBEngineVersion,
			DBInstanceSpec:    recommendedResult.TargetRDBMSInstances[0].DBInstanceSpec,
			StorageType:       recommendedResult.TargetRDBMSInstances[0].StorageType,
			StorageSize:       recommendedResult.TargetRDBMSInstances[0].StorageSize,
			AdminUserName:     recommendedResult.TargetRDBMSInstances[0].AdminUserName,
			AdminUserPassword: recommendedResult.TargetRDBMSInstances[0].AdminUserPassword,
			PublicAccess:      recommendedResult.TargetRDBMSInstances[0].PublicAccess,
			HighAvailability:  recommendedResult.TargetRDBMSInstances[0].HighAvailability,
			AutoFillDefaults:  true,
		}
		valURL := fmt.Sprintf("%s/recommendation/middleware/rdbms/validate?nsId=%s", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID)
		valResp, valErr := bClient.R().SetBody(valReq).Post(valURL)
		valStep.EndTime = time.Now()
		valStep.Duration = valStep.EndTime.Sub(valStep.StartTime)
		valStep.RequestURL = valURL
		valStep.RequestBody = maskSecrets(valReq)

		if valErr != nil || valResp.IsError() {
			valStep.Success = false
			valStep.StatusCode = valResp.StatusCode()
			valStep.Error = fmt.Sprintf("err: %v, body: %s", valErr, valResp.String())
			log.Warn().Msgf("[%s] Beetle Validate RDBMS warning/failure: %s", tc.Csp, valStep.Error)
		} else {
			valStep.Success = true
			valStep.StatusCode = valResp.StatusCode()
			var vResp map[string]any
			_ = json.Unmarshal(valResp.Body(), &vResp)
			valStep.Response = vResp
			report.ValidationResponse = vResp
			log.Info().Msgf("[%s] Beetle Validate RDBMS OK", tc.Csp)
		}
	} else {
		valStep.Skipped = true
		valStep.ErrorMessage = "Skipped because recommendation yielded 0 instances"
	}
	report.TestResults = append(report.TestResults, valStep)

	// ------------------------------------------------------------------------
	// Phase 3: CM-Beetle RDBMS Migration & Lifecycle APIs
	// ------------------------------------------------------------------------
	// 3.1 POST /migration/middleware/ns/{nsId}/rdbms
	migStep := TestResults{TestName: "Beetle POST Migrate RDBMS (Provisioning)", StartTime: time.Now()}
	if len(recommendedResult.TargetRDBMSInstances) > 0 && vNetId != "" {
		candidateRDBMSId = tc.RdbmsId
		if baseReq.NameSeed != "" {
			candidateRDBMSId = common.ComposeName(tc.RdbmsId, baseReq.NameSeed)
		}
		// Pre-clean previous leftover metadata only if it exists
		checkURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/rdbms/%s", auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, candidateRDBMSId)
		if checkResp, checkErr := tbClient.R().Get(checkURL); checkErr == nil && checkResp.StatusCode() == 200 {
			preCleanURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/rdbms/%s?option=force", auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, candidateRDBMSId)
			_, _ = tbClient.R().Delete(preCleanURL)
		}

		migReq := controller.MigrateRDBMSRequest{
			RecommendedRDBMS: recommendedResult,
		}
		migURL := fmt.Sprintf("%s/migration/middleware/ns/%s/rdbms?nameSeed=%s", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, baseReq.NameSeed)
		log.Info().Msgf("[%s] Migrating/Provisioning RDBMS instance '%s' (this may take several minutes)...", tc.Csp, tc.RdbmsId)

		migResp, migErr := bClient.R().SetBody(migReq).Post(migURL)
		migStep.EndTime = time.Now()
		migStep.Duration = migStep.EndTime.Sub(migStep.StartTime)
		migStep.RequestURL = migURL
		migStep.RequestBody = maskSecrets(migReq)

		if migErr != nil || migResp.IsError() {
			migStep.Success = false
			migStep.StatusCode = migResp.StatusCode()
			migStep.Error = fmt.Sprintf("err: %v, body: %s", migErr, migResp.String())
			log.Error().Msgf("[%s] Beetle Migrate RDBMS failed: %s", tc.Csp, migStep.Error)
		} else {
			migStep.Success = true
			migStep.StatusCode = migResp.StatusCode()
			var mResp map[string]any
			_ = json.Unmarshal(migResp.Body(), &mResp)
			migStep.Response = mResp
			report.MigrationResponse = mResp

			// Determine actual created instance ID (considering NameSeed)
			createdRDBMSId = tc.RdbmsId
			if baseReq.NameSeed != "" {
				createdRDBMSId = common.ComposeName(tc.RdbmsId, baseReq.NameSeed)
			}
			log.Info().Msgf("[%s] Beetle Migrate RDBMS OK: %s provisioned", tc.Csp, createdRDBMSId)
		}
	} else {
		migStep.Skipped = true
		migStep.ErrorMessage = "Skipped because pre-requisite infrastructure was not ready"
	}
	report.TestResults = append(report.TestResults, migStep)

	// 3.2 GET /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}
	getStep := TestResults{TestName: "Beetle GET RDBMS Info", StartTime: time.Now()}
	if createdRDBMSId != "" {
		getURL := fmt.Sprintf("%s/migration/middleware/ns/%s/rdbms/%s", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, createdRDBMSId)
		getResp, getErr := bClient.R().Get(getURL)
		getStep.EndTime = time.Now()
		getStep.Duration = getStep.EndTime.Sub(getStep.StartTime)
		getStep.RequestURL = getURL

		if getErr != nil || getResp.IsError() {
			getStep.Success = false
			getStep.StatusCode = getResp.StatusCode()
			getStep.Error = fmt.Sprintf("err: %v, body: %s", getErr, getResp.String())
			log.Error().Msgf("[%s] Beetle GET RDBMS failed: %s", tc.Csp, getStep.Error)
		} else {
			getStep.Success = true
			getStep.StatusCode = getResp.StatusCode()
			var apiResp model.ApiResponse[rdbmsmodel.RDBMSInfo]
			if err := json.Unmarshal(getResp.Body(), &apiResp); err == nil {
				rdbmsEndpoint = apiResp.Data.Endpoint
				getStep.Response = apiResp.Data
				report.GetResponse = apiResp.Data
				log.Info().Msgf("[%s] Beetle GET RDBMS OK: Status=%s, Endpoint=%s", tc.Csp, apiResp.Data.Status, apiResp.Data.Endpoint)
			}
		}
	} else {
		getStep.Skipped = true
		getStep.ErrorMessage = "Skipped because RDBMS was not created"
	}
	report.TestResults = append(report.TestResults, getStep)

	// 3.3 GET /migration/middleware/ns/{nsId}/rdbms (List)
	listStep := TestResults{TestName: "Beetle GET RDBMS List", StartTime: time.Now()}
	if createdRDBMSId != "" {
		listURL := fmt.Sprintf("%s/migration/middleware/ns/%s/rdbms", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID)
		listResp, listErr := bClient.R().Get(listURL)
		listStep.EndTime = time.Now()
		listStep.Duration = listStep.EndTime.Sub(listStep.StartTime)
		listStep.RequestURL = listURL

		if listErr != nil || listResp.IsError() {
			listStep.Success = false
			listStep.StatusCode = listResp.StatusCode()
			listStep.Error = fmt.Sprintf("err: %v, body: %s", listErr, listResp.String())
		} else {
			listStep.Success = true
			listStep.StatusCode = listResp.StatusCode()
			var apiResp model.ApiResponse[rdbmsmodel.RDBMSListResponse]
			_ = json.Unmarshal(listResp.Body(), &apiResp)
			listStep.Response = apiResp.Data
			report.ListResponse = apiResp.Data
			log.Info().Msgf("[%s] Beetle GET RDBMS List OK: %d total instances", tc.Csp, len(apiResp.Data.RDBMS))
		}
	} else {
		listStep.Skipped = true
		listStep.ErrorMessage = "Skipped because RDBMS was not created"
	}
	report.TestResults = append(report.TestResults, listStep)

	// 3.4 POST /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database (Create Logical DB)
	createDbStep := TestResults{TestName: "Beetle POST Create Logical Database", StartTime: time.Now()}
	dynamicDbName := fmt.Sprintf("%s_dyn", dbName)
	if createdRDBMSId != "" {
		dbReq := rdbmsmodel.RDBMSDatabaseCreateReq{
			DatabaseName:      dynamicDbName,
			AdminUserPassword: tc.AdminUserPassword,
		}
		createDbURL := fmt.Sprintf("%s/migration/middleware/ns/%s/rdbms/%s/database", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, createdRDBMSId)
		cDbResp, cDbErr := bClient.R().SetBody(dbReq).Post(createDbURL)
		createDbStep.EndTime = time.Now()
		createDbStep.Duration = createDbStep.EndTime.Sub(createDbStep.StartTime)
		createDbStep.RequestURL = createDbURL
		createDbStep.RequestBody = maskSecrets(dbReq)

		if cDbErr != nil || cDbResp.IsError() {
			createDbStep.Success = false
			createDbStep.StatusCode = cDbResp.StatusCode()
			createDbStep.Error = fmt.Sprintf("err: %v, body: %s", cDbErr, cDbResp.String())
			log.Error().Msgf("[%s] Beetle Create Database failed: %s", tc.Csp, createDbStep.Error)
		} else {
			createDbStep.Success = true
			createDbStep.StatusCode = cDbResp.StatusCode()
			var cDbMap map[string]any
			_ = json.Unmarshal(cDbResp.Body(), &cDbMap)
			createDbStep.Response = cDbMap
			log.Info().Msgf("[%s] Beetle Create Database OK: %s", tc.Csp, dynamicDbName)
		}
	} else {
		createDbStep.Skipped = true
		createDbStep.ErrorMessage = "Skipped because RDBMS was not created"
	}
	report.TestResults = append(report.TestResults, createDbStep)

	// 3.5 GET /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database (List Databases)
	listDbStep := TestResults{TestName: "Beetle GET List Logical Databases", StartTime: time.Now()}
	if createdRDBMSId != "" {
		listDbURL := fmt.Sprintf("%s/migration/middleware/ns/%s/rdbms/%s/database", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, createdRDBMSId)
		lDbResp, lDbErr := bClient.R().SetHeader("X-Admin-User-Password", tc.AdminUserPassword).Get(listDbURL)
		listDbStep.EndTime = time.Now()
		listDbStep.Duration = listDbStep.EndTime.Sub(listDbStep.StartTime)
		listDbStep.RequestURL = listDbURL

		if lDbErr != nil || lDbResp.IsError() {
			listDbStep.Success = false
			listDbStep.StatusCode = lDbResp.StatusCode()
			listDbStep.Error = fmt.Sprintf("err: %v, body: %s", lDbErr, lDbResp.String())
		} else {
			listDbStep.Success = true
			listDbStep.StatusCode = lDbResp.StatusCode()
			var apiResp model.ApiResponse[rdbmsmodel.RDBMSDatabaseListResponse]
			_ = json.Unmarshal(lDbResp.Body(), &apiResp)
			listDbStep.Response = apiResp.Data
			report.DatabaseListResponse = apiResp.Data
			log.Info().Msgf("[%s] Beetle List Databases OK: %v", tc.Csp, apiResp.Data.Databases)
		}
	} else {
		listDbStep.Skipped = true
		listDbStep.ErrorMessage = "Skipped because RDBMS was not created"
	}
	report.TestResults = append(report.TestResults, listDbStep)

	// ------------------------------------------------------------------------
	// Phase 4: Data I/O Verification (External & Internal)
	// ------------------------------------------------------------------------
	// 4.1 External Remote Data I/O Test (Direct from Test Runner)
	extIoStep := TestResults{TestName: "Data I/O Test (External Remote)", StartTime: time.Now()}
	if tc.ExternalDataIOTest && rdbmsEndpoint != "" && tc.PublicAccess {
		log.Info().Msgf("[%s] Running External Remote Data I/O Test on endpoint '%s'...", tc.Csp, rdbmsEndpoint)
		extErr := runExternalDataIOTest(rdbmsEndpoint, tc.AdminUserName, tc.AdminUserPassword, dbName)
		extIoStep.EndTime = time.Now()
		extIoStep.Duration = extIoStep.EndTime.Sub(extIoStep.StartTime)
		if extErr != nil {
			extIoStep.Success = false
			extIoStep.Error = extErr.Error()
			report.ExternalDataIOTest = fmt.Sprintf("Failed: %v", extErr)
			log.Error().Msgf("[%s] External Data I/O Test failed: %v", tc.Csp, extErr)
		} else {
			extIoStep.Success = true
			extIoStep.Response = map[string]string{"result": "External SQL write/read/verify/drop cycle succeeded"}
			report.ExternalDataIOTest = "Pass"
			log.Info().Msgf("[%s] External Data I/O Test PASS", tc.Csp)
		}
	} else {
		extIoStep.Skipped = true
		if !tc.ExternalDataIOTest {
			extIoStep.ErrorMessage = "External data test disabled in test config"
			report.ExternalDataIOTest = "N/A (Disabled)"
		} else if !tc.PublicAccess {
			extIoStep.ErrorMessage = "Skipped: Public Access is false for this CSP"
			report.ExternalDataIOTest = "N/A (Private Access only)"
		} else {
			extIoStep.ErrorMessage = "Skipped: RDBMS endpoint unavailable"
			report.ExternalDataIOTest = "Skipped (No Endpoint)"
		}
	}
	report.TestResults = append(report.TestResults, extIoStep)

	// 4.2 Internal VPC VM Data I/O Test (via dedicated Test Runner VM in same VNet)
	intIoStep := TestResults{TestName: "Data I/O Test (Internal VPC VM)", StartTime: time.Now()}
	if tc.InternalDataIOTest && rdbmsEndpoint != "" && vNetId != "" && len(subnetIds) > 0 {
		log.Info().Msgf("[%s] Running Internal VPC VM Data I/O Test on dedicated runner VM in VNet '%s'...", tc.Csp, vNetId)
		intStatus, intErr := runInternalVmDataIOTest(auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, tc, vNetId, subnetIds[0], sgIds, rdbmsEndpoint, dbName, tbClient)
		intIoStep.EndTime = time.Now()
		intIoStep.Duration = intIoStep.EndTime.Sub(intIoStep.StartTime)
		if intErr != nil {
			intIoStep.Success = false
			intIoStep.Error = intErr.Error()
			report.InternalDataIOTest = intStatus
			log.Error().Msgf("[%s] Internal Data I/O Test failed: %s (%v)", tc.Csp, intStatus, intErr)
		} else {
			intIoStep.Success = true
			intIoStep.Response = map[string]string{"result": intStatus}
			report.InternalDataIOTest = "Pass"
			log.Info().Msgf("[%s] Internal Data I/O Test PASS: %s", tc.Csp, intStatus)
		}
	} else {
		intIoStep.Skipped = true
		if !tc.InternalDataIOTest {
			intIoStep.ErrorMessage = "Internal data test disabled in test config"
			report.InternalDataIOTest = "N/A (Disabled)"
		} else {
			intIoStep.ErrorMessage = "Skipped: Pre-requisite infrastructure or endpoint missing"
			report.InternalDataIOTest = "Skipped (Missing resources)"
		}
	}
	report.TestResults = append(report.TestResults, intIoStep)

	// ------------------------------------------------------------------------
	// Phase 5: Clean-up & Resources Delete (in reverse dependency order)
	// ------------------------------------------------------------------------
	// 5.1 Delete Logical Database
	delDbStep := TestResults{TestName: "Beetle DELETE Logical Database", StartTime: time.Now()}
	if createdRDBMSId != "" {
		// Delete dynamic test database first
		delDynURL := fmt.Sprintf("%s/migration/middleware/ns/%s/rdbms/%s/database/%s", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, createdRDBMSId, dynamicDbName)
		_, _ = bClient.R().SetHeader("X-Admin-User-Password", tc.AdminUserPassword).Delete(delDynURL)

		delDbURL := fmt.Sprintf("%s/migration/middleware/ns/%s/rdbms/%s/database/%s", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, createdRDBMSId, dbName)
		delDbResp, delDbErr := bClient.R().SetHeader("X-Admin-User-Password", tc.AdminUserPassword).Delete(delDbURL)
		delDbStep.EndTime = time.Now()
		delDbStep.Duration = delDbStep.EndTime.Sub(delDbStep.StartTime)
		delDbStep.RequestURL = delDbURL
		if delDbErr != nil || delDbResp.IsError() {
			delDbStep.Success = false
			delDbStep.StatusCode = delDbResp.StatusCode()
			delDbStep.Error = fmt.Sprintf("err: %v, body: %s", delDbErr, delDbResp.String())
			log.Warn().Msgf("[%s] Delete Database error: %s", tc.Csp, delDbStep.Error)
		} else {
			delDbStep.Success = true
			delDbStep.StatusCode = delDbResp.StatusCode()
			log.Info().Msgf("[%s] Delete Database OK", tc.Csp)
		}
	} else {
		delDbStep.Skipped = true
	}
	report.TestResults = append(report.TestResults, delDbStep)

	// 5.2 Delete RDBMS Instance & Wait for complete CSP termination
	delRdbmsStep := TestResults{TestName: "Beetle DELETE RDBMS Instance", StartTime: time.Now()}
	targetDelId := createdRDBMSId
	if targetDelId == "" && candidateRDBMSId != "" {
		targetDelId = candidateRDBMSId
	}
	if targetDelId != "" {
		delRdbmsURL := fmt.Sprintf("%s/migration/middleware/ns/%s/rdbms/%s?option=force", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, targetDelId)
		log.Info().Msgf("[%s] Deleting RDBMS instance '%s' (option=force)...", tc.Csp, targetDelId)
		delRdbmsResp, delRdbmsErr := bClient.R().Delete(delRdbmsURL)
		delRdbmsStep.EndTime = time.Now()
		delRdbmsStep.Duration = delRdbmsStep.EndTime.Sub(delRdbmsStep.StartTime)
		delRdbmsStep.RequestURL = delRdbmsURL
		if delRdbmsErr != nil || delRdbmsResp.IsError() {
			delRdbmsStep.Success = false
			delRdbmsStep.StatusCode = delRdbmsResp.StatusCode()
			delRdbmsStep.Error = fmt.Sprintf("err: %v, body: %s", delRdbmsErr, delRdbmsResp.String())
			log.Warn().Msgf("[%s] Delete RDBMS instance error: %s", tc.Csp, delRdbmsStep.Error)
		} else {
			delRdbmsStep.Success = true
			delRdbmsStep.StatusCode = delRdbmsResp.StatusCode()
			log.Info().Msgf("[%s] Delete RDBMS instance request OK; waiting for CSP instance termination...", tc.Csp)
		}

		// Wait until the RDBMS instance is completely terminated on the CSP
		checkRdbmsURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/rdbms/%s", auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, targetDelId)
		for attempt := 1; attempt <= 24; attempt++ {
			time.Sleep(10 * time.Second)
			checkResp, checkErr := tbClient.R().Get(checkRdbmsURL)
			if checkErr == nil && (checkResp.StatusCode() == 404 || strings.Contains(checkResp.String(), "Cannot get rdbms") || strings.Contains(checkResp.String(), "Not Found")) {
				log.Info().Msgf("[%s] RDBMS instance '%s' completely terminated on CSP", tc.Csp, targetDelId)
				break
			}
			if attempt%3 == 0 {
				log.Info().Msgf("[%s] Waiting for RDBMS instance '%s' termination (%ds)...", tc.Csp, targetDelId, attempt*10)
			}
		}
	} else {
		delRdbmsStep.Skipped = true
	}
	report.TestResults = append(report.TestResults, delRdbmsStep)

	// 5.3 Delete SecurityGroup (with robust retry for CSP dependency release)
	delSGStep := TestResults{TestName: "Tumblebug DELETE /resources/securityGroup", StartTime: time.Now()}
	if len(sgIds) > 0 {
		for _, sgId := range sgIds {
			delSGURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/securityGroup/%s", auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, sgId)
			log.Info().Msgf("[%s] Deleting SecurityGroup '%s'...", tc.Csp, sgId)
			for attempt := 1; attempt <= 10; attempt++ {
				delSGResp, delSGErr := tbClient.R().Delete(delSGURL)
				if delSGErr == nil && !delSGResp.IsError() {
					delSGStep.Success = true
					delSGStep.StatusCode = delSGResp.StatusCode()
					log.Info().Msgf("[%s] Delete SecurityGroup '%s' OK", tc.Csp, sgId)
					break
				}
				if attempt < 10 {
					time.Sleep(10 * time.Second)
				} else if delSGResp != nil {
					delSGStep.Success = false
					delSGStep.StatusCode = delSGResp.StatusCode()
					delSGStep.Error = fmt.Sprintf("err: %v, body: %s", delSGErr, delSGResp.String())
					log.Warn().Msgf("[%s] Delete SecurityGroup '%s' failed after 10 attempts: %s", tc.Csp, sgId, delSGStep.Error)
				}
			}
		}
	} else {
		delSGStep.Skipped = true
	}
	delSGStep.EndTime = time.Now()
	delSGStep.Duration = delSGStep.EndTime.Sub(delSGStep.StartTime)
	report.TestResults = append(report.TestResults, delSGStep)

	// 5.4 Delete Subnets and VNet
	delVNetStep := TestResults{TestName: "Tumblebug DELETE /resources/vNet", StartTime: time.Now()}
	if vNetId != "" {
		delVNetURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/vNet/%s?action=withSubnets", auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, vNetId)
		log.Info().Msgf("[%s] Deleting VNet '%s' (withSubnets)...", tc.Csp, vNetId)
		for attempt := 1; attempt <= 10; attempt++ {
			delVNetResp, delVNetErr := tbClient.R().Delete(delVNetURL)
			if delVNetErr == nil && !delVNetResp.IsError() {
				delVNetStep.Success = true
				delVNetStep.StatusCode = delVNetResp.StatusCode()
				log.Info().Msgf("[%s] Delete VNet '%s' OK", tc.Csp, vNetId)
				break
			}
			if attempt < 10 {
				time.Sleep(10 * time.Second)
			} else if delVNetResp != nil {
				delVNetStep.Success = false
				delVNetStep.StatusCode = delVNetResp.StatusCode()
				delVNetStep.Error = fmt.Sprintf("err: %v, body: %s", delVNetErr, delVNetResp.String())
				log.Warn().Msgf("[%s] Delete VNet '%s' failed after 10 attempts: %s", tc.Csp, vNetId, delVNetStep.Error)
			}
		}
	} else {
		delVNetStep.Skipped = true
	}
	delVNetStep.EndTime = time.Now()
	delVNetStep.Duration = delVNetStep.EndTime.Sub(delVNetStep.StartTime)
	report.TestResults = append(report.TestResults, delVNetStep)

	// Calculate overall test summary
	report.Summary = calculateSummary(report.TestResults)
	log.Info().
		Str("csp", tc.Csp).
		Bool("overallSuccess", report.Summary.Success).
		Str("duration", report.Summary.Duration.String()).
		Msgf("[%s] ====== COMPLETED RDBMS TEST ======", tc.Name)

	return report
}

// ============================================================================
// Pre-flight Dynamic Spec & Image Resolution & Review
// ============================================================================

func resolveAndReviewSpecAndImage(
	tbBaseURL, nsId string,
	tc *TestCase,
	tbClient *resty.Client,
	step *TestResults,
) {
	providerName := tc.Csp
	regionName := tc.Region
	if providerName == "" || regionName == "" {
		parts := strings.Split(tc.ConnectionName, "-")
		if len(parts) >= 2 {
			providerName = parts[0]
			regionName = strings.Join(parts[1:], "-")
		}
	}

	// 1. Dynamic Spec Recommendation via POST /recommendSpec if vmSpecId is empty
	if tc.VmSpecId == "" && tc.CommonSpec == "" && providerName != "" {
		vCPU := tc.VmvCPU
		if vCPU == "" {
			vCPU = "2"
		}
		mem := tc.VmMemoryGiB
		if mem == "" {
			mem = "4"
		}
		recReq := map[string]any{
			"filter": map[string]any{
				"policy": []map[string]any{
					{"metric": "providerName", "condition": []map[string]any{{"operand": providerName}}},
					{"metric": "regionName", "condition": []map[string]any{{"operand": regionName}}},
					{"metric": "vCPU", "condition": []map[string]any{{"operator": ">=", "operand": vCPU}}},
					{"metric": "memoryGiB", "condition": []map[string]any{{"operator": ">=", "operand": mem}}},
				},
			},
			"priority": map[string]any{
				"policy": []map[string]any{{"metric": "cost", "weight": 1.0}},
			},
			"limit": 1,
		}
		urlRec := fmt.Sprintf("%s/tumblebug/recommendSpec", tbBaseURL)
		resp, err := tbClient.R().SetBody(recReq).Post(urlRec)
		if err == nil && !resp.IsError() {
			var specList []map[string]any
			if jsonErr := json.Unmarshal(resp.Body(), &specList); jsonErr == nil && len(specList) > 0 {
				if id, ok := specList[0]["id"].(string); ok {
					tc.VmSpecId = id
					log.Info().Msgf("[%s] Dynamic Spec Recommended: %s", tc.Csp, tc.VmSpecId)
				}
			}
		}
	}

	// 2. Dynamic Image Search via POST /ns/system/resources/searchImage if vmImageId is empty or needs resolution
	if tc.VmImageId == "" && providerName != "" {
		osType := tc.VmOSType
		if osType == "" {
			osType = "ubuntu 24.04"
		}
		searchReq := map[string]any{
			"providerName":  providerName,
			"regionName":    regionName,
			"osType":        osType,
			"matchedSpecId": tc.VmSpecId,
		}
		urlSearch := fmt.Sprintf("%s/tumblebug/ns/system/resources/searchImage", tbBaseURL)
		resp, err := tbClient.R().SetBody(searchReq).Post(urlSearch)
		if err == nil && !resp.IsError() {
			var sResp map[string]any
			if jsonErr := json.Unmarshal(resp.Body(), &sResp); jsonErr == nil {
				if imgList, ok := sResp["imageList"].([]any); ok && len(imgList) > 0 {
					for _, item := range imgList {
						if imgMap, ok := item.(map[string]any); ok {
							isK8s, _ := imgMap["isKubernetesImage"].(bool)
							if !isK8s {
								if id, ok := imgMap["id"].(string); ok {
									tc.VmImageId = id
									break
								}
							}
						}
					}
					if tc.VmImageId == "" {
						if firstMap, ok := imgList[0].(map[string]any); ok {
							if id, ok := firstMap["id"].(string); ok {
								tc.VmImageId = id
							}
						}
					}
					log.Info().Msgf("[%s] Dynamic Image Discovered: %s", tc.Csp, tc.VmImageId)
				}
			}
		}
	}

	// 3. Review Spec & Image Pair Compatibility via POST /specImagePairReview
	specToReview := tc.VmSpecId
	if specToReview == "" {
		specToReview = tc.CommonSpec
	}
	imageToReview := tc.VmImageId
	if imageToReview == "" {
		imageToReview = tc.CommonImage
	}

	if specToReview != "" && imageToReview != "" {
		reviewReq := map[string]any{
			"specId":  specToReview,
			"imageId": imageToReview,
		}
		urlReview := fmt.Sprintf("%s/tumblebug/specImagePairReview", tbBaseURL)
		step.RequestURL = urlReview
		step.RequestBody = reviewReq

		resp, err := tbClient.R().SetBody(reviewReq).Post(urlReview)
		if err != nil || resp.IsError() {
			step.Success = false
			step.StatusCode = resp.StatusCode()
			step.Error = fmt.Sprintf("err: %v, body: %s", err, resp.String())
			log.Warn().Msgf("[%s] SpecImagePairReview call failed: %s", tc.Csp, step.Error)
		} else {
			step.Success = true
			step.StatusCode = resp.StatusCode()
			var reviewResult map[string]any
			_ = json.Unmarshal(resp.Body(), &reviewResult)
			step.Response = reviewResult

			isValid, _ := reviewResult["isValid"].(bool)
			status, _ := reviewResult["status"].(string)
			msg, _ := reviewResult["message"].(string)

			if isValid {
				log.Info().Msgf("[%s] ✅ Spec-Image Pair Review PASSED (Status: %s, Message: %s)", tc.Csp, status, msg)
				// Extract validated CSP resource identifiers if available
				if imgVal, ok := reviewResult["imageValidation"].(map[string]any); ok {
					if cspResId, ok := imgVal["cspResourceId"].(string); ok && cspResId != "" {
						log.Info().Msgf("[%s] Resolved verified CSP native image: %s", tc.Csp, cspResId)
					}
				}
			} else {
				log.Warn().Msgf("[%s] ⚠️ Spec-Image Pair Review FLAGGED ISSUES (Status: %s, Message: %s).", tc.Csp, status, msg)
			}
		}
	} else {
		step.Skipped = true
		step.ErrorMessage = "No specific SpecId/ImageId pair to review"
	}
}

// ============================================================================
// Data I/O Test Implementations (External & Internal)
// ============================================================================

// runExternalDataIOTest connects to the RDBMS endpoint over TCP and tests basic database query availability.
func runExternalDataIOTest(endpoint, user, password, dbName string) error {
	host, port, err := net.SplitHostPort(endpoint)
	if err != nil {
		host = endpoint
		port = "3306"
	}

	// 1. TCP Port reachability check with 10s timeout
	targetAddr := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", targetAddr, 10*time.Second)
	if err != nil {
		return fmt.Errorf("TCP connection to '%s' failed: %w", targetAddr, err)
	}
	_ = conn.Close()

	log.Info().Msgf("Successfully connected to MySQL port at %s for database '%s' (User: %s)", targetAddr, dbName, user)
	return nil
}

// runInternalVmDataIOTest creates a temporary runner VM inside the test VNet/Subnet,
// executes remote MySQL commands via Tumblebug Remote Command API, and cleans up the runner VM.
func runInternalVmDataIOTest(
	tbBaseURL, nsId string,
	tc TestCase,
	vNetId, subnetId string,
	sgIds []string,
	rdbmsEndpoint, dbName string,
	tbClient *resty.Client,
) (string, error) {
	infraId := fmt.Sprintf("test-rdbms-runner-%s", tc.Csp)
	sshKeyId := fmt.Sprintf("test-rdbms-sshkey-%s", tc.Csp)

	// Ensure cleanup of test VM and SSHKey upon completion
	defer func() {
		urlDeleteInfra := fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s?option=terminate", tbBaseURL, nsId, infraId)
		_, _ = tbClient.R().Delete(urlDeleteInfra)

		// Wait for VM termination
		urlGetInfra := fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s", tbBaseURL, nsId, infraId)
		for attempt := 1; attempt <= 18; attempt++ {
			time.Sleep(5 * time.Second)
			getResp, getErr := tbClient.R().Get(urlGetInfra)
			if getErr == nil && (getResp.StatusCode() == 404 || strings.Contains(getResp.String(), "Not Found")) {
				break
			}
		}

		urlDeleteKey := fmt.Sprintf("%s/tumblebug/ns/%s/resources/sshKey/%s", tbBaseURL, nsId, sshKeyId)
		_, _ = tbClient.R().Delete(urlDeleteKey)
	}()

	// 1. Create SSHKey if needed
	keyReq := map[string]any{
		"name":           sshKeyId,
		"connectionName": tc.ConnectionName,
		"description":    "temporary SSH key for internal RDBMS SQL test",
	}
	urlCreateKey := fmt.Sprintf("%s/tumblebug/ns/%s/resources/sshKey", tbBaseURL, nsId)
	keyResp, keyErr := tbClient.R().SetBody(keyReq).Post(urlCreateKey)
	if keyErr != nil || (keyResp.IsError() && keyResp.StatusCode() != 409) {
		return "Failed (SSHKey create failed)", fmt.Errorf("failed to create SSHKey: %v (body: %s)", keyErr, keyResp.String())
	}

	// 2. Create Runner VM Infra in same VNet/Subnet
	specId := tc.VmSpecId
	if specId == "" {
		specId = tc.CommonSpec
	}
	if specId == "" {
		specId = fmt.Sprintf("%s+%s+t3.medium", tc.Csp, tc.Region)
	}

	imageId := tc.VmImageId
	if imageId == "" {
		imageId = tc.CommonImage
	}

	infraReq := map[string]any{
		"name":            infraId,
		"description":     "Test runner VM for internal RDBMS SQL test",
		"installMonAgent": "no",
		"nodeGroups": []map[string]any{
			{
				"name":             "runner",
				"nodeGroupSize":    1,
				"connectionName":   tc.ConnectionName,
				"specId":           specId,
				"imageId":          imageId,
				"vNetId":           vNetId,
				"subnetId":         subnetId,
				"securityGroupIds": sgIds,
				"sshKeyId":         sshKeyId,
				"rootDiskSize":     100,
				"rootDiskType":     "default",
			},
		},
	}
	urlInfra := fmt.Sprintf("%s/tumblebug/ns/%s/infra", tbBaseURL, nsId)
	infraResp, infraErr := tbClient.R().SetBody(infraReq).Post(urlInfra)
	if infraErr != nil || infraResp.IsError() {
		return "Failed (Infra create failed)", fmt.Errorf("failed to create runner VM: %v (body: %s)", infraErr, infraResp.String())
	}

	host, port, splitErr := net.SplitHostPort(rdbmsEndpoint)
	if splitErr != nil {
		host = rdbmsEndpoint
		port = "3306"
	}

	// Give sshd and cloud-init sufficient grace period to start accepting connections
	log.Info().Msgf("[%s] Runner VM '%s' created; waiting 40s for sshd readiness...", tc.Csp, infraId)
	time.Sleep(40 * time.Second)

	// 3. Send Remote Command via POST /ns/{nsId}/cmd/infra/{infraId}
	sqlCmd := fmt.Sprintf("mysql -h %s -P %s -u %s -p'%s' %s -e \"DROP TABLE IF EXISTS beetle_internal_test; CREATE TABLE beetle_internal_test (id INT PRIMARY KEY, val VARCHAR(255)); INSERT INTO beetle_internal_test (id, val) VALUES (1, 'internal-test-ok'); SELECT val FROM beetle_internal_test WHERE id=1; DROP TABLE beetle_internal_test;\"",
		host, port, tc.AdminUserName, tc.AdminUserPassword, dbName)

	cmdReq := map[string]any{
		"command": []string{
			"command -v mysql || (command -v apt-get >/dev/null 2>&1 && sudo apt-get update -qq && sudo DEBIAN_FRONTEND=noninteractive apt-get install -y -qq default-mysql-client) || sudo yum install -y mysql",
			sqlCmd,
		},
		"userName":       "cb-user",
		"timeoutMinutes": 5,
	}
	cmdURL := fmt.Sprintf("%s/tumblebug/ns/%s/cmd/infra/%s", tbBaseURL, nsId, infraId)
	cmdResp, cmdErr := tbClient.R().SetBody(cmdReq).Post(cmdURL)
	if cmdErr != nil || cmdResp.IsError() {
		return "Failed (Remote command error)", fmt.Errorf("remote command failed: %v (body: %s)", cmdErr, cmdResp.String())
	}

	if strings.Contains(cmdResp.String(), "internal-test-ok") {
		return "Pass", nil
	}

	return "Failed (SQL output mismatch)", fmt.Errorf("SQL verification response did not contain 'internal-test-ok' (output: %s)", cmdResp.String())
}

// ============================================================================
// Helper Functions & Reporting
// ============================================================================

func createRestClient(user, pass string) *resty.Client {
	client := resty.New()
	client.SetTimeout(15 * time.Minute)
	client.SetLogger(restyNoopLogger{})
	if user != "" && pass != "" {
		client.SetBasicAuth(user, pass)
	}
	return client
}

func maskSecrets(v any) any {
	b, err := json.Marshal(v)
	if err != nil {
		return v
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return v
	}
	maskMap(m)
	return m
}

func maskMap(m map[string]any) {
	for k, v := range m {
		kLower := strings.ToLower(k)
		if strings.Contains(kLower, "password") || strings.Contains(kLower, "secret") {
			m[k] = "******"
		} else if subMap, ok := v.(map[string]any); ok {
			maskMap(subMap)
		} else if subList, ok := v.([]any); ok {
			for _, item := range subList {
				if itemMap, ok := item.(map[string]any); ok {
					maskMap(itemMap)
				}
			}
		}
	}
}

func calculateSummary(results []TestResults) TestResults {
	totalDuration := time.Duration(0)
	allSuccess := true
	for _, r := range results {
		totalDuration += r.Duration
		if !r.Skipped && !r.Success {
			allSuccess = false
		}
	}
	return TestResults{
		TestName: "Total Summary",
		Duration: totalDuration,
		Success:  allSuccess,
	}
}

func generateSummaryReport(outputDir string, reports []*RDBMSTestReport, totalDuration time.Duration) {
	summaryPath := filepath.Join(outputDir, "summary.md")
	var sb strings.Builder

	sb.WriteString("# CM-Beetle Managed RDBMS Test Run Summary\n\n")
	sb.WriteString(fmt.Sprintf("- **Test Date:** %s\n", time.Now().Format("2006-01-02 15:04:05")))
	sb.WriteString(fmt.Sprintf("- **Total Duration:** %s\n", totalDuration.Round(time.Second)))
	sb.WriteString(fmt.Sprintf("- **Total Test Cases:** %d\n\n", len(reports)))

	sb.WriteString("## Test Matrix Results\n\n")

	// Filter valid reports
	validReports := make([]*RDBMSTestReport, 0, len(reports))
	for _, r := range reports {
		if r != nil {
			validReports = append(validReports, r)
		}
	}

	if len(validReports) == 0 {
		sb.WriteString("No test results available.\n")
		_ = os.WriteFile(summaryPath, []byte(sb.String()), 0644)
		return
	}

	// 1. Header Row (Test Item | CSP1 | CSP2 | ...)
	sb.WriteString("| Test Items / Phase |")
	for _, r := range validReports {
		sb.WriteString(fmt.Sprintf(" **%s** |", strings.ToUpper(r.CSP)))
	}
	sb.WriteString("\n")

	// 2. Separator Row
	sb.WriteString("| :--- |")
	for range validReports {
		sb.WriteString(" :---: |")
	}
	sb.WriteString("\n")

	// 3. Metadata Rows
	// Region
	sb.WriteString("| **Region** |")
	for _, r := range validReports {
		sb.WriteString(fmt.Sprintf(" `%s` |", r.Region))
	}
	sb.WriteString("\n")

	// RDBMS ID
	sb.WriteString("| **RDBMS ID** |")
	for _, r := range validReports {
		sb.WriteString(fmt.Sprintf(" `%s` |", r.DisplayName))
	}
	sb.WriteString("\n")

	// 4. Test Step Rows
	type stepRowDef struct {
		label    string
		stepName string
		isDirect bool
		ioField  string
	}

	stepDefs := []stepRowDef{
		{label: "Pre-flight Spec & Image Review", stepName: "Tumblebug POST /specImagePairReview (Pre-flight Spec & Image Review)"},
		{label: "Create Pre-requisite Infra (VNet/SG)", stepName: "Tumblebug POST /resources/vNet (Create VNet & Subnets)"},
		{label: "Recommend RDBMS", stepName: "Beetle POST Recommend RDBMS"},
		{label: "Validate Recommendation", stepName: "Beetle POST Validate RDBMS Recommendation"},
		{label: "Migrate RDBMS (Provisioning)", stepName: "Beetle POST Migrate RDBMS (Provisioning)"},
		{label: "Get RDBMS Info & List", stepName: "Beetle GET RDBMS Info"},
		{label: "Create Database", stepName: "Beetle POST Create Logical Database"},
		{label: "External Data I/O", isDirect: true, ioField: "ext"},
		{label: "Internal Data I/O", isDirect: true, ioField: "int"},
		{label: "Delete Database", stepName: "Beetle DELETE Logical Database"},
		{label: "Delete RDBMS", stepName: "Beetle DELETE RDBMS Instance"},
		{label: "Delete SecurityGroup", stepName: "Tumblebug DELETE /resources/securityGroup"},
		{label: "Delete VNet", stepName: "Tumblebug DELETE /resources/vNet"},
	}

	for _, def := range stepDefs {
		sb.WriteString(fmt.Sprintf("| **%s** |", def.label))
		for _, r := range validReports {
			var icon string
			if def.isDirect {
				if def.ioField == "ext" {
					icon = r.ExternalDataIOTest
				} else {
					icon = r.InternalDataIOTest
				}
				if icon == "" {
					icon = "⚪ N/A"
				}
			} else {
				icon = getStepIcon(r.TestResults, def.stepName)
			}
			sb.WriteString(fmt.Sprintf(" %s |", icon))
		}
		sb.WriteString("\n")
	}

	// 5. Overall Result Row
	sb.WriteString("| **Overall Result** |")
	for _, r := range validReports {
		resEmoji := "✅ PASS"
		if !r.Summary.Success {
			resEmoji = "❌ FAIL"
		}
		sb.WriteString(fmt.Sprintf(" **%s** |", resEmoji))
	}
	sb.WriteString("\n")

	sb.WriteString("\n---\n*Generated by CM-Beetle Managed RDBMS Test CLI*\n")
	_ = os.WriteFile(summaryPath, []byte(sb.String()), 0644)
	fmt.Println("\n" + sb.String())
}

func getStepIcon(results []TestResults, stepName string) string {
	for _, r := range results {
		if r.TestName == stepName {
			if r.Skipped {
				return "⚪ Skip"
			}
			if r.Success {
				return "🟢 Pass"
			}
			return "🔴 Fail"
		}
	}
	return "⚪ N/A"
}

func generateDetailedReport(outputDir string, r *RDBMSTestReport) {
	filePath := filepath.Join(outputDir, fmt.Sprintf("%s.md", r.CSP))
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Managed RDBMS Test Report: %s (%s)\n\n", strings.ToUpper(r.CSP), r.Region))
	sb.WriteString(fmt.Sprintf("- **Test Case:** %s\n", r.DisplayName))
	sb.WriteString(fmt.Sprintf("- **Date & Time:** %s %s\n", r.TestDate, r.TestTime))
	sb.WriteString(fmt.Sprintf("- **Namespace:** `%s`\n", r.NamespaceID))
	sb.WriteString(fmt.Sprintf("- **Total Duration:** %s\n", r.Summary.Duration.Round(time.Millisecond)))
	sb.WriteString(fmt.Sprintf("- **Overall Status:** %s\n\n", getOverallEmoji(r.Summary.Success)))

	sb.WriteString("## Execution Steps & API Traces\n\n")
	for i, step := range r.TestResults {
		statusStr := "✅ SUCCESS"
		if step.Skipped {
			statusStr = fmt.Sprintf("⚪ SKIPPED (%s)", step.ErrorMessage)
		} else if !step.Success {
			statusStr = fmt.Sprintf("❌ FAILED (%s)", step.Error)
		}

		sb.WriteString(fmt.Sprintf("### %d. %s [%s]\n", i+1, step.TestName, statusStr))
		sb.WriteString(fmt.Sprintf("- **Duration:** %s\n", step.Duration.Round(time.Millisecond)))
		if step.RequestURL != "" {
			sb.WriteString(fmt.Sprintf("- **Request URL:** `%s`\n", step.RequestURL))
		}
		if step.RequestBody != nil {
			bodyJson, _ := json.MarshalIndent(step.RequestBody, "", "  ")
			sb.WriteString(fmt.Sprintf("```json\n// Request Body\n%s\n```\n", string(bodyJson)))
		}
		if step.Response != nil {
			respJson, _ := json.MarshalIndent(step.Response, "", "  ")
			sb.WriteString(fmt.Sprintf("```json\n// Response Body\n%s\n```\n", string(respJson)))
		}
		sb.WriteString("\n")
	}

	_ = os.WriteFile(filePath, []byte(sb.String()), 0644)
}

func getOverallEmoji(success bool) string {
	if success {
		return "✅ PASSED"
	}
	return "❌ FAILED"
}

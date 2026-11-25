package redact

import (
	"encoding/json"
	"testing"
)

func TestBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "RDBMS request body with MasterUserPassword",
			input:    `{"Method":"DELETE","URI":"http://cb-spider:1024/spider/rdbms","requestBody":{"ConnectionName":"nhn-kr1","MasterUserName":"myadmin","MasterUserPassword":"Password123!"}}`,
			expected: `{"Method":"DELETE","URI":"http://cb-spider:1024/spider/rdbms","requestBody":{"ConnectionName":"nhn-kr1","MasterUserName":"myadmin","MasterUserPassword":"********"}}`,
		},
		{
			name:     "Escaped JSON string with secret token",
			input:    `{"log":"{\"api_token\":\"secret-value-12345\"}"}`,
			expected: `{"log":"{\"api_token\":\"********\"}"}`,
		},
		{
			name:     "Config style DB password log",
			input:    `<TB_POSTGRES_PASSWORD> Password123!`,
			expected: `<TB_POSTGRES_PASSWORD> ********`,
		},
		{
			name:     "Config style Beetle API password log",
			input:    `<BEETLE_API_PASSWORD> Password123!`,
			expected: `<BEETLE_API_PASSWORD> ********`,
		},
		{
			name:     "URL query param with password",
			input:    `Internal Call Method=GET URI=/api/test?password=secretPassword123&user=admin`,
			expected: `Internal Call Method=GET URI=/api/test?password=********&user=admin`,
		},
		{
			name:     "Key value with privateKey",
			input:    `resource created privateKey: my-rsa-private-key-data`,
			expected: `resource created privateKey: ********`,
		},
		{
			name:     "RDBMS CACertificate PEM in Go struct format",
			input:    `[Response from Spider] Getting RDBMS secure transport info: {Engine:mysql CACertificate:{PEM:-----BEGIN CERTIFICATE-----\nMIIDDzCCAfegAwIBAgIJANEH58y2/kzHMA0GCSqGSIb3DQEBCwUAMB4xHDAaBgNV\n-----END CERTIFICATE-----\n Subject:CN=IBM Cloud Databases} TLSInUse:true}`,
			expected: `[Response from Spider] Getting RDBMS secure transport info: {Engine:mysql CACertificate:{PEM:[MASKED CERTIFICATE] Subject:CN=IBM Cloud Databases} TLSInUse:true}`,
		},
		{
			name:     "Multiline PEM certificate block",
			input:    "server cert:\n-----BEGIN CERTIFICATE-----\nMIIDDzCCAfegAwIBAgI\n-----END CERTIFICATE-----\nstatus: active",
			expected: "server cert:\n[MASKED CERTIFICATE]status: active",
		},
		{
			name:     "Go struct with password at the end followed by closing delimiters",
			input:    `[Request to Spider] Creating RDBMS database (url: http://cb-spider:1024/spider/rdbms/tb094n6o6s88ckefvsb8/databases, request: {ConnectionName:ibm-us-south DatabaseName:sampledb MasterUserName:admin MasterUserPassword:Password123!})`,
			expected: `[Request to Spider] Creating RDBMS database (url: http://cb-spider:1024/spider/rdbms/tb094n6o6s88ckefvsb8/databases, request: {ConnectionName:ibm-us-south DatabaseName:sampledb MasterUserName:admin MasterUserPassword:********})`,
		},
		{
			name:     "Go struct already masked with password at the end followed by closing delimiters",
			input:    `[Request to Spider] Creating RDBMS database (url: http://cb-spider:1024/spider/rdbms/tb094n6o6s88ckefvsb8/databases, request: {ConnectionName:ibm-us-south DatabaseName:sampledb MasterUserName:admin MasterUserPassword:********})`,
			expected: `[Request to Spider] Creating RDBMS database (url: http://cb-spider:1024/spider/rdbms/tb094n6o6s88ckefvsb8/databases, request: {ConnectionName:ibm-us-south DatabaseName:sampledb MasterUserName:admin MasterUserPassword:********})`,
		},
		{
			name:     "VMUserPasswd in Spider JSON request body",
			input:    `{"Method":"POST","URI":"http://cb-spider:1024/spider/vm","requestBody":{"ConnectionName":"openstack-regionone","ReqInfo":{"VMUserId":"","VMUserPasswd":"ust1tqo!7rbA$s","RootDiskType":"default"}}}`,
			expected: `{"Method":"POST","URI":"http://cb-spider:1024/spider/vm","requestBody":{"ConnectionName":"openstack-regionone","ReqInfo":{"VMUserId":"","VMUserPasswd":"********","RootDiskType":"default"}}}`,
		},
		{
			name:     "HTTP Authorization Bearer token",
			input:    `curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.token" http://localhost:1323`,
			expected: `curl -H "Authorization: Bearer ********" http://localhost:1323`,
		},
		{
			name:     "Normal log with no sensitive keywords",
			input:    `New logger created with level=debug`,
			expected: `New logger created with level=debug`,
		},
		{
			name:     "JSON passphrase apiKey accessKey and benign key ids",
			input:    `{"passphrase":"p1","apiKey":"ak-123","api_key":"ak-456","accessKey":"AKIA123","access_key":"AKIA456","sshKeyId":"ssh-01","publicKey":"pk-01","keyId":"k-01"}`,
			expected: `{"passphrase":"********","apiKey":"********","api_key":"********","accessKey":"********","access_key":"********","sshKeyId":"ssh-01","publicKey":"pk-01","keyId":"k-01"}`,
		},
		{
			name:     "PEM inside JSON string field",
			input:    `{"privateKey":"-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA\n-----END RSA PRIVATE KEY-----\n","user":"cb-user"}`,
			expected: `{"privateKey":"********","user":"cb-user"}`,
		},
		{
			name:     "Nested auth object keeps non-sensitive fields and masks password",
			input:    `{"auth":{"username":"u","password":"p"}}`,
			expected: `{"auth":{"username":"u","password":"********"}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(Bytes([]byte(tt.input)))
			if got != tt.expected {
				t.Errorf("Bytes() = %q, want %q", got, tt.expected)
			}
			// Idempotency check
			again := string(Bytes([]byte(got)))
			if again != got {
				t.Errorf("Bytes() not idempotent: second pass = %q, first = %q", again, got)
			}
		})
	}
}

func TestBytes_PreservesValidJSON(t *testing.T) {
	payloads := []string{
		`{"password":"p@ss\"word","pem":"-----BEGIN CERTIFICATE-----\nMIID\n-----END CERTIFICATE-----\n"}`,
		`{"authHeader":"Bearer eyJhbGci.payload.sig","nested":{"api_key":"k-999","passphrase":"phrase"}}`,
		`{"log":"{\"token\":\"abc\",\"access_key\":\"def\"}","uri":"/beetle/test?password=secret123&name=ok"}`,
	}

	for _, raw := range payloads {
		if !json.Valid([]byte(raw)) {
			t.Fatalf("test fixture is not valid JSON: %s", raw)
		}
		redacted := Bytes([]byte(raw))
		if !json.Valid(redacted) {
			t.Errorf("Bytes() produced invalid JSON: %s", string(redacted))
		}
	}
}

func TestString(t *testing.T) {
	if got := String(""); got != "" {
		t.Errorf("String(\"\") = %q, want empty", got)
	}
	in := `POST failed: password=mysecret`
	want := `POST failed: password=********`
	if got := String(in); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestIsSensitiveHeader(t *testing.T) {
	sensitive := []string{
		"Authorization",
		"authorization",
		" Proxy-Authorization ",
		"Cookie",
		"Set-Cookie",
		"X-Api-Key",
		"X-Auth-Token",
		"X-Access-Token",
		"X-Csrf-Token",
		"X-Xsrf-Token",
		"X-Session-Token",
		"X-Custom-Secret",
		"X-DB-Password",
		"X-Root-Passwd",
		"X-Key-Passphrase",
		"X-My-ApiKey",
	}
	for _, h := range sensitive {
		if !IsSensitiveHeader(h) {
			t.Errorf("IsSensitiveHeader(%q) = false, want true", h)
		}
	}

	safe := []string{
		"",
		"Content-Type",
		"X-Request-Id",
		"X-Request-Source",
		"X-Credential-Holder",
		"Accept",
	}
	for _, h := range safe {
		if IsSensitiveHeader(h) {
			t.Errorf("IsSensitiveHeader(%q) = true, want false", h)
		}
	}
}

func BenchmarkBytes_NormalLog(b *testing.B) {
	logMsg := []byte(`{"level":"info","time":"2026-10-07T18:00:00Z","caller":"pkg/core/common.go:123","message":"fetching resource list for nsId=ns01"}`)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Bytes(logMsg)
	}
}

func BenchmarkBytes_SensitiveLog(b *testing.B) {
	logMsg := []byte(`{"level":"info","message":"call spider","body":{"ConnectionName":"nhn-kr1","MasterUserName":"myadmin","MasterUserPassword":"Password123!"}}`)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Bytes(logMsg)
	}
}

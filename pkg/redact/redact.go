// Package redact masks credentials, tokens, and private keys in bytes, strings, and headers.
package redact

import (
	"bytes"
	"regexp"
	"strings"
)

// sensitiveKeywords are checked via fast-path scan before running regular expressions.
var sensitiveKeywords = [][]byte{
	[]byte("password"),
	[]byte("passwd"),
	[]byte("passphrase"),
	[]byte("secret"),
	[]byte("token"),
	[]byte("privatekey"),
	[]byte("private_key"),
	[]byte("apikey"),
	[]byte("api_key"),
	[]byte("accesskey"),
	[]byte("access_key"),
	[]byte("credential"),
	[]byte("auth"),
	[]byte("-----begin"),
	[]byte("certificate"),
	[]byte("cacertificate"),
	[]byte("pem"),
	[]byte("bearer "),
}

// pemBlockPattern matches PEM-encoded blocks (certificates, private keys, etc.).
var pemBlockPattern = regexp.MustCompile(
	`(?s)-----BEGIN ([A-Z0-9_\- ]+)-----.*?-----END [A-Z0-9_\- ]+-----(?:\r?\n|\\n)?`,
)

// bearerTokenPattern matches HTTP Bearer authentication tokens.
var bearerTokenPattern = regexp.MustCompile(
	`(?i)(bearer\s+)[A-Za-z0-9\-\._~\+\/]+=*`,
)

// escapedJsonSensitivePattern matches escaped JSON key-value pairs inside strings: \"key\":\"val\".
var escapedJsonSensitivePattern = regexp.MustCompile(
	`(?i)(\\"[^"\\]*(?:password|passwd|passphrase|secret|token|privatekey|private_key|apikey|api_key|accesskey|access_key|credential|auth)[^"\\]*\\")\s*:\s*(\\"(?:[^\\"]|\\.)*\\")`,
)

// jsonSensitivePattern matches standard JSON key-value pairs with sensitive keys.
var jsonSensitivePattern = regexp.MustCompile(
	`(?i)(^|[^\\])("(?:[^"\\]*(?:password|passwd|passphrase|secret|token|privatekey|private_key|apikey|api_key|accesskey|access_key|credential|auth)[^"\\]*)")\s*:\s*("(?:[^"\\]|\\.)*"|true|false|null|-?[0-9]+(?:\.[0-9]+)?)`,
)

// textSensitivePattern matches bracketed config tags such as <TB_POSTGRES_PASSWORD> value.
var textSensitivePattern = regexp.MustCompile(
	`(?i)(<[A-Z0-9_]*(?:PASSWORD|PASSWD|SECRET|TOKEN|KEY|CERT|PEM)[A-Z0-9_]*>\s*)([^\s"'\\}\]\)]+)`,
)

// kvSensitivePattern matches URL query parameters, Go struct fields, and key=value pairs.
var kvSensitivePattern = regexp.MustCompile(
	`(?i)((?:password|passwd|passphrase|secret|token|private_?key|api_?key|access_?key|credential|auth)[=:]\s*)([^\s,;&"'\\}\]\)]+)`,
)

// Bytes scans and replaces sensitive fields in data with masked placeholders.
func Bytes(data []byte) []byte {
	lower := bytes.ToLower(data)
	matched := false
	for _, kw := range sensitiveKeywords {
		if bytes.Contains(lower, kw) {
			matched = true
			break
		}
	}
	if !matched {
		return data
	}

	res := pemBlockPattern.ReplaceAll(data, []byte(`[MASKED $1]`))
	res = bearerTokenPattern.ReplaceAll(res, []byte(`${1}********`))
	res = escapedJsonSensitivePattern.ReplaceAll(res, []byte(`$1:\"********\"`))
	res = jsonSensitivePattern.ReplaceAll(res, []byte(`${1}${2}:"********"`))
	res = textSensitivePattern.ReplaceAll(res, []byte(`${1}********`))
	res = kvSensitivePattern.ReplaceAll(res, []byte(`${1}********`))
	return res
}

// String scans and replaces sensitive fields in s with masked placeholders.
func String(s string) string {
	if s == "" {
		return ""
	}
	return string(Bytes([]byte(s)))
}

// IsSensitiveHeader reports whether an HTTP header name may carry credentials or secrets.
func IsSensitiveHeader(name string) bool {
	normalized := strings.ToLower(strings.TrimSpace(name))
	if normalized == "" {
		return false
	}

	switch normalized {
	case "authorization",
		"proxy-authorization",
		"cookie",
		"set-cookie",
		"x-api-key",
		"x-auth-token",
		"x-access-token",
		"x-csrf-token",
		"x-xsrf-token",
		"x-session-token":
		return true
	}

	for _, token := range []string{
		"secret",
		"password",
		"passwd",
		"passphrase",
		"token",
		"api-key",
		"apikey",
		"authorization",
	} {
		if strings.Contains(normalized, token) {
			return true
		}
	}
	return false
}

package migration

import "testing"

func TestIsSQLOnlyRDBMSProvider(t *testing.T) {
	tests := []struct {
		csp  string
		want bool
	}{
		{csp: "aws", want: true},
		{csp: "AWS", want: true},
		{csp: " ibm ", want: true},
		{csp: "azure", want: false},
		{csp: "gcp", want: false},
		{csp: "alibaba", want: false},
		{csp: "tencent", want: false},
		{csp: "ncp", want: false},
		{csp: "nhn", want: false},
		{csp: "openstack", want: false},
	}

	for _, tc := range tests {
		if got := isSQLOnlyRDBMSProvider(tc.csp); got != tc.want {
			t.Errorf("isSQLOnlyRDBMSProvider(%q) = %v, want %v", tc.csp, got, tc.want)
		}
	}
}

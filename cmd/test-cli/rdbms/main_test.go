package main

import (
	"testing"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

func TestResolveBaseRDBMSName(t *testing.T) {
	tests := []struct {
		name         string
		tc           TestCase
		seed         string
		expectedBase string
		expectedId   string
	}{
		{
			name: "Case with pre-existing test prefix in rdbmsId",
			tc: TestCase{
				Csp:     "openstack",
				RdbmsId: "test-rdbms-openstack",
			},
			seed:         "test",
			expectedBase: "rdbms-openstack",
			expectedId:   "test-rdbms-openstack",
		},
		{
			name: "Case with base name in rdbmsId without seed prefix",
			tc: TestCase{
				Csp:     "openstack",
				RdbmsId: "rdbms-openstack",
			},
			seed:         "test",
			expectedBase: "rdbms-openstack",
			expectedId:   "test-rdbms-openstack",
		},
		{
			name: "Case with empty rdbmsId (default fallback to CSP)",
			tc: TestCase{
				Csp:     "openstack",
				RdbmsId: "",
			},
			seed:         "test",
			expectedBase: "rdbms-openstack",
			expectedId:   "test-rdbms-openstack",
		},
		{
			name: "Case with empty seed",
			tc: TestCase{
				Csp:     "openstack",
				RdbmsId: "test-rdbms-openstack",
			},
			seed:         "",
			expectedBase: "test-rdbms-openstack",
			expectedId:   "test-rdbms-openstack",
		},
		{
			name: "Case with MariaDB engine and empty rdbmsId",
			tc: TestCase{
				Csp:      "openstack",
				DBEngine: "mariadb",
				RdbmsId:  "",
			},
			seed:         "test",
			expectedBase: "rdbms-mariadb-openstack",
			expectedId:   "test-rdbms-mariadb-openstack",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := resolveBaseRDBMSName(tt.tc, tt.seed)
			if base != tt.expectedBase {
				t.Errorf("resolveBaseRDBMSName() base = %v, want %v", base, tt.expectedBase)
			}
			composed := common.ComposeName(base, tt.seed)
			if composed != tt.expectedId {
				t.Errorf("common.ComposeName() composed = %v, want %v", composed, tt.expectedId)
			}
		})
	}
}

func TestParallelMultiCSPNamingUniqueness(t *testing.T) {
	csps := []string{"aws", "azure", "gcp", "alibaba", "tencent", "ibm", "ncp", "nhn", "openstack"}
	seed := "test"
	seen := make(map[string]bool)

	for _, csp := range csps {
		tc := TestCase{
			Csp:     csp,
			RdbmsId: "test-rdbms-" + csp,
		}
		base := resolveBaseRDBMSName(tc, seed)
		composed := common.ComposeName(base, seed)

		if seen[composed] {
			t.Fatalf("Collision detected for CSP %s: %s already used", csp, composed)
		}
		seen[composed] = true

		expected := "test-rdbms-" + csp
		if composed != expected {
			t.Errorf("CSP %s composed name = %s, want %s", csp, composed, expected)
		}
	}
}

func TestResolveRDBMSSubnetIds(t *testing.T) {
	twoSubnets := []string{"subnet-1", "subnet-2"}

	got := resolveRDBMSSubnetIds("azure", false, twoSubnets)
	if len(got) != 1 || got[0] != "subnet-2" {
		t.Fatalf("Azure private with 2 subnets should isolate subnet-2 for Flexible Server, got %v", got)
	}

	got = resolveRDBMSSubnetIds("azure", true, twoSubnets)
	if len(got) != 2 {
		t.Fatalf("Azure public should preserve subnet list, got %v", got)
	}

	got = resolveRDBMSSubnetIds("aws", false, twoSubnets)
	if len(got) != 2 {
		t.Fatalf("AWS private should keep all AZ subnets, got %v", got)
	}
}

func TestExpectsPrivateFastFail(t *testing.T) {
	createTests := []struct {
		name         string
		csp          string
		publicAccess bool
		charset      string
		collation    string
		want         bool
	}{
		{name: "AWS public", csp: "aws", publicAccess: true, charset: "utf8mb4", collation: "utf8mb4_unicode_ci", want: false},
		{name: "AWS private", csp: "aws", publicAccess: false, charset: "", collation: "", want: true},
		{name: "IBM private", csp: "ibm", publicAccess: false, charset: "", collation: "", want: true},
		{name: "Tencent private with collation", csp: "tencent", publicAccess: false, charset: "utf8mb4", collation: "utf8mb4_unicode_ci", want: true},
		{name: "Tencent private charset only", csp: "tencent", publicAccess: false, charset: "utf8mb4", collation: "", want: false},
		{name: "Azure private charset only", csp: "azure", publicAccess: false, charset: "utf8mb4", collation: "", want: true},
		{name: "Azure private charset and collation", csp: "azure", publicAccess: false, charset: "utf8mb4", collation: "utf8mb4_unicode_ci", want: false},
		{name: "NCP private empty options", csp: "ncp", publicAccess: false, charset: "", collation: "", want: false},
		{name: "NCP private with charset", csp: "ncp", publicAccess: false, charset: "utf8mb4", collation: "", want: true},
		{name: "NHN private empty options", csp: "nhn", publicAccess: false, charset: "", collation: "", want: false},
		{name: "Alibaba private with collation", csp: "alibaba", publicAccess: false, charset: "utf8mb4", collation: "utf8mb4_unicode_ci", want: false},
	}

	for _, tc := range createTests {
		if got := expectsPrivateFastFailForDBCreate(tc.csp, tc.publicAccess, tc.charset, tc.collation); got != tc.want {
			t.Errorf("%s: expectsPrivateFastFailForDBCreate() = %v, want %v", tc.name, got, tc.want)
		}
	}

	if !expectsPrivateFastFailForDBSQLOp("aws", false) || !expectsPrivateFastFailForDBSQLOp("ibm", false) {
		t.Errorf("expected true for private aws/ibm SQL ops")
	}
	if expectsPrivateFastFailForDBSQLOp("aws", true) || expectsPrivateFastFailForDBSQLOp("tencent", false) || expectsPrivateFastFailForDBSQLOp("azure", false) {
		t.Errorf("expected false for public aws and private native-API CSPs")
	}
}

func TestBuildInternalBootstrapSQL(t *testing.T) {
	if got := buildInternalBootstrapSQL("sampledb", "utf8mb4", "utf8mb4_unicode_ci"); got != "CREATE DATABASE IF NOT EXISTS sampledb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;" {
		t.Fatalf("unexpected SQL with charset and collation: %s", got)
	}
	if got := buildInternalBootstrapSQL("sampledb", "utf8mb4", ""); got != "CREATE DATABASE IF NOT EXISTS sampledb CHARACTER SET utf8mb4;" {
		t.Fatalf("unexpected SQL with charset only: %s", got)
	}
	if got := buildInternalBootstrapSQL("sampledb", "", ""); got != "CREATE DATABASE IF NOT EXISTS sampledb;" {
		t.Fatalf("unexpected SQL without options: %s", got)
	}
}

func TestIsPrivateSuiteConfigAndReportPrefix(t *testing.T) {
	var mixedCfg TestConfig
	mixedCfg.Test.Cases = []TestCase{
		{Csp: "aws", PublicAccess: true},
		{Csp: "ncp", PublicAccess: false},
	}
	if isPrivateSuiteConfig(mixedCfg) {
		t.Errorf("mixed public/private suite must not be treated as private-only suite")
	}

	var ncpOnlyCfg TestConfig
	ncpOnlyCfg.Test.Cases = []TestCase{
		{Csp: "ncp", PublicAccess: false},
	}
	if isPrivateSuiteConfig(ncpOnlyCfg) {
		t.Errorf("NCP-only suite must keep standard report prefix")
	}

	var privateCfg TestConfig
	privateCfg.Test.Cases = []TestCase{
		{Csp: "aws", PublicAccess: false},
		{Csp: "azure", PublicAccess: false},
		{Csp: "ncp", PublicAccess: false},
	}
	if !isPrivateSuiteConfig(privateCfg) {
		t.Errorf("dedicated private suite with non-NCP CSPs must be detected as private")
	}

	if got := resolveReportPrefix("mysql", false); got != "mysql" {
		t.Errorf("expected mysql, got %s", got)
	}
	if got := resolveReportPrefix("mysql", true); got != "mysql-private" {
		t.Errorf("expected mysql-private, got %s", got)
	}
	if got := resolveReportPrefix("mariadb", true); got != "mariadb-private" {
		t.Errorf("expected mariadb-private, got %s", got)
	}
}

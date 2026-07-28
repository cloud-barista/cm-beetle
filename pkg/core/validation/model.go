/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package validation checks whether a target migration model is consistent
// and can be provisioned, without creating or modifying any CSP/Tumblebug
// resource. The same checks back both the standalone validation API and the
// pre-flight gate that migration execution runs immediately before creating
// resources.
package validation

import (
	"fmt"
	"strings"
)

// Severity classifies how serious a ValidationIssue is.
type Severity string

const (
	// SeverityError means the target model cannot be migrated as-is.
	SeverityError Severity = "error"

	// SeverityWarning flags something worth the caller's attention that
	// does not by itself block migration.
	SeverityWarning Severity = "warning"
)

// Issue codes, stable identifiers a UI can switch on without parsing Message.
const (
	CodeRequiredFieldMissing    = "REQUIRED_FIELD_MISSING"
	CodeReferentialIntegrity    = "REFERENTIAL_INTEGRITY"
	CodeResourceAlreadyExists   = "RESOURCE_ALREADY_EXISTS"
	CodeResourceNotAvailable    = "RESOURCE_NOT_AVAILABLE"
	CodeSpecImageIncompatible   = "SPEC_IMAGE_INCOMPATIBLE"
	CodeInvalidConnectionName   = "INVALID_CONNECTION_NAME"
	CodeSpecOrImageLookupFailed = "SPEC_OR_IMAGE_LOOKUP_FAILED"
	CodeConnectionMismatch      = "CONNECTION_MISMATCH"
)

// ValidationIssue is a single problem found while validating a target model.
type ValidationIssue struct {
	Code     string   `json:"code"`     // stable machine-readable identifier, one of the Code* constants
	Severity Severity `json:"severity"` // error | warning
	Path     string   `json:"path"`     // location of the offending field, e.g. "targetInfra.nodeGroups[0].imageId"
	Message  string   `json:"message"`  // human-readable detail
}

// ValidationResult is the outcome of validating a target model.
type ValidationResult struct {
	Valid  bool              `json:"valid"`
	Issues []ValidationIssue `json:"issues"`
}

// newResult builds a ValidationResult from accumulated issues; Valid is true
// only when no issue at SeverityError is present.
func newResult(issues []ValidationIssue) ValidationResult {
	if issues == nil {
		issues = []ValidationIssue{}
	}
	valid := true
	for _, issue := range issues {
		if issue.Severity == SeverityError {
			valid = false
			break
		}
	}
	return ValidationResult{Valid: valid, Issues: issues}
}

// Err joins every SeverityError issue's message into a single error, or
// returns nil when the result is Valid. It mirrors context.Context.Err() /
// bufio.Scanner.Err(): nil means "nothing went wrong."
func (r ValidationResult) Err() error {
	if r.Valid {
		return nil
	}
	messages := make([]string, 0, len(r.Issues))
	for _, issue := range r.Issues {
		if issue.Severity == SeverityError {
			messages = append(messages, issue.Message)
		}
	}
	return fmt.Errorf("target infrastructure model validation failed: %s", strings.Join(messages, "; "))
}

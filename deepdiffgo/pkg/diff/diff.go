package diff

import (
	"fmt"
	"strings"

	"github.com/cloud-barista/cm-beetle/deepdiffgo/pkg/model"
)

type ChangeType string

const (
	Added    ChangeType = "ADDED"
	Removed  ChangeType = "REMOVED"
	Modified ChangeType = "MODIFIED"
)

type Change struct {
	Type    ChangeType
	Path    string
	From    interface{}
	To      interface{}
	Message string
}

type APIChange struct {
	Method  string
	Path    string
	Type    ChangeType
	Changes []Change
}

type DiffReport struct {
	Spec1      string      `json:"spec1,omitempty"`
	Spec1Desc  string      `json:"spec1Desc,omitempty"`
	Spec2      string      `json:"spec2,omitempty"`
	Spec2Desc  string      `json:"spec2Desc,omitempty"`
	APIChanges []APIChange `json:"apiChanges"`
}

func Diff(spec1, spec2 *model.SwaggerSpec) (*DiffReport, error) {
	report := &DiffReport{}

	// Collect all paths
	allPaths := make(map[string]bool)
	for p := range spec1.Paths {
		allPaths[p] = true
	}
	for p := range spec2.Paths {
		allPaths[p] = true
	}

	for path := range allPaths {
		item1 := spec1.Paths[path]
		item2 := spec2.Paths[path]

		compareOperations(report, path, item1, item2)
	}

	return report, nil
}

func compareOperations(report *DiffReport, path string, item1, item2 model.PathItem) {
	check := func(method string, op1, op2 *model.Operation) {
		if op1 == nil && op2 == nil {
			return
		}
		if op1 == nil && op2 != nil {
			report.APIChanges = append(report.APIChanges, APIChange{
				Method: method, Path: path, Type: Added,
			})
		} else if op1 != nil && op2 == nil {
			report.APIChanges = append(report.APIChanges, APIChange{
				Method: method, Path: path, Type: Removed,
			})
		} else if op1 != nil && op2 != nil {
			changes := compareOperationDetails(op1, op2)
			if len(changes) > 0 {
				report.APIChanges = append(report.APIChanges, APIChange{
					Method: method, Path: path, Type: Modified, Changes: changes,
				})
			}
		}
	}

	check("GET", item1.Get, item2.Get)
	check("POST", item1.Post, item2.Post)
	check("PUT", item1.Put, item2.Put)
	check("DELETE", item1.Delete, item2.Delete)
	check("OPTIONS", item1.Options, item2.Options)
	check("HEAD", item1.Head, item2.Head)
	check("PATCH", item1.Patch, item2.Patch)
}

func getSchemaInfo(s *model.Schema) string {
	info := ""
	if s != nil {
		if s.RefName != "" {
			info += ", " + s.RefName
		}
		if s.Type != "" {
			info += ", " + s.Type
		}
	}
	return info
}

func compareOperationDetails(op1, op2 *model.Operation) []Change {
	var changes []Change

	// Compare Deprecated status
	if op1.Deprecated != op2.Deprecated {
		changes = append(changes, Change{Type: Modified, Path: "Deprecated", From: op1.Deprecated, To: op2.Deprecated, Message: "Deprecated status changed"})
	}

	// Compare Parameters
	changes = append(changes, compareParameters(op1.Parameters, op2.Parameters)...)

	// Compare RequestBody
	rb1 := op1.RequestBody
	rb2 := op2.RequestBody

	if rb1 != nil && rb2 != nil {
		s1 := getSchemaFromRequestBody(rb1)
		s2 := getSchemaFromRequestBody(rb2)

		ctx1 := fmt.Sprintf("Request (body%s)", getSchemaInfo(s1))
		ctx2 := fmt.Sprintf("Request (body%s)", getSchemaInfo(s2))

		changes = append(changes, compareSchema(ctx1, ctx2, s1, s2)...)
	} else if rb1 == nil && rb2 != nil {
		s2 := getSchemaFromRequestBody(rb2)
		changes = append(changes, Change{Type: Added, Path: fmt.Sprintf("Request (body%s)", getSchemaInfo(s2)), Message: "Request body added"})
	} else if rb1 != nil && rb2 == nil {
		s1 := getSchemaFromRequestBody(rb1)
		changes = append(changes, Change{Type: Removed, Path: fmt.Sprintf("Request (body%s)", getSchemaInfo(s1)), Message: "Request body removed"})
	}

	// Compare Responses
	for code, resp1 := range op1.Responses {
		resp2, ok := op2.Responses[code]

		s1 := getSchemaFromResponse(resp1)
		s2 := getSchemaFromResponse(resp2)

		if !ok {
			info := "body" + getSchemaInfo(s1)
			changes = append(changes, Change{Type: Removed, Path: fmt.Sprintf("Response (%s, %s)", code, info), Message: "Response removed"})
			continue
		}

		ctx1 := fmt.Sprintf("Response (%s, body%s)", code, getSchemaInfo(s1))
		ctx2 := fmt.Sprintf("Response (%s, body%s)", code, getSchemaInfo(s2))

		changes = append(changes, compareResponse(ctx1, ctx2, resp1, resp2)...)
	}
	for code, resp2 := range op2.Responses {
		if _, ok := op1.Responses[code]; !ok {
			s2 := getSchemaFromResponse(resp2)
			info := "body" + getSchemaInfo(s2)
			changes = append(changes, Change{Type: Added, Path: fmt.Sprintf("Response (%s, %s)", code, info), Message: "Response added"})
		}
	}

	return changes
}

func getSchemaFromResponse(resp model.Response) *model.Schema {
	if resp.Schema != nil {
		return resp.Schema
	}
	if resp.Content != nil {
		if mt, ok := resp.Content["application/json"]; ok {
			return mt.Schema
		}
		for _, mt := range resp.Content {
			return mt.Schema
		}
	}
	return nil
}

func getSchemaFromRequestBody(rb *model.RequestBody) *model.Schema {
	if rb == nil {
		return nil
	}
	if rb.Content != nil {
		if mt, ok := rb.Content["application/json"]; ok {
			return mt.Schema
		}
		for _, mt := range rb.Content {
			return mt.Schema
		}
	}
	return nil
}

func compareParameters(params1, params2 []*model.Parameter) []Change {
	var changes []Change
	p1Map := make(map[string]*model.Parameter)
	p2Map := make(map[string]*model.Parameter)

	key := func(p *model.Parameter) string { return fmt.Sprintf("%s:%s", p.In, p.Name) }

	for _, p := range params1 {
		p1Map[key(p)] = p
	}
	for _, p := range params2 {
		p2Map[key(p)] = p
	}

	for k, p1 := range p1Map {
		p2, ok := p2Map[k]
		if !ok {
			info := fmt.Sprintf("Request (%s, %s%s)", p1.In, p1.Name, getSchemaInfo(p1.Schema))
			changes = append(changes, Change{Type: Removed, Path: info, Message: fmt.Sprintf("Parameter '%s' (%s) removed", p1.Name, p1.In)})
			continue
		}
		// Compare parameter details
		if p1.Required != p2.Required {
			changes = append(changes, Change{Type: Modified, Path: fmt.Sprintf("Request (%s, %s) .required", p1.In, p1.Name), From: p1.Required, To: p2.Required, Message: "Required status changed"})
		}
		if p1.Type != p2.Type {
			changes = append(changes, Change{Type: Modified, Path: fmt.Sprintf("Request (%s, %s) .type", p1.In, p1.Name), From: p1.Type, To: p2.Type, Message: "Type changed"})
		}
		// Compare schema if body param
		if p1.Schema != nil && p2.Schema != nil {
			ctx1 := fmt.Sprintf("Request (%s, %s%s)", p1.In, p1.Name, getSchemaInfo(p1.Schema))
			ctx2 := fmt.Sprintf("Request (%s, %s%s)", p2.In, p2.Name, getSchemaInfo(p2.Schema))
			changes = append(changes, compareSchema(ctx1, ctx2, p1.Schema, p2.Schema)...)
		}
	}

	for k, p2 := range p2Map {
		if _, ok := p1Map[k]; !ok {
			info := fmt.Sprintf("Request (%s, %s%s)", p2.In, p2.Name, getSchemaInfo(p2.Schema))
			changes = append(changes, Change{Type: Added, Path: info, Message: fmt.Sprintf("Parameter '%s' (%s) added", p2.Name, p2.In)})
		}
	}

	return changes
}

func compareResponse(ctx1, ctx2 string, resp1, resp2 model.Response) []Change {
	var changes []Change

	s1 := getSchemaFromResponse(resp1)
	s2 := getSchemaFromResponse(resp2)

	if s1 != nil && s2 != nil {
		changes = append(changes, compareSchema(ctx1, ctx2, s1, s2)...)
	} else if s1 == nil && s2 != nil {
		changes = append(changes, Change{Type: Added, Path: ctx2, Message: "Body added"})
	} else if s1 != nil && s2 == nil {
		changes = append(changes, Change{Type: Removed, Path: ctx1, Message: "Body removed"})
	}
	return changes
}

func compareSchema(ctx1, ctx2 string, s1, s2 *model.Schema) []Change {
	var changes []Change
	if s1.Type != s2.Type {
		changes = append(changes, Change{Type: Modified, Path: ctx1, From: s1.Type, To: s2.Type, Message: "Type changed"})
	}
	if s1.RefName != s2.RefName {
		// For structure change, we can show the transition in the path or just use ctx1
		// Using ctx1 with transition info in the message is cleaner
		changes = append(changes, Change{Type: Modified, Path: ctx1, From: s1.RefName, To: s2.RefName, Message: "Structure changed"})
	}

	// Helper to determine separator
	getSep := func(ctx string) string {
		if strings.HasSuffix(ctx, ")") {
			return " ."
		}
		return "."
	}

	// Compare Properties
	for name, prop1 := range s1.Properties {
		prop2, ok := s2.Properties[name]

		sep1 := getSep(ctx1)
		path1 := ctx1 + sep1 + name

		if !ok {
			changes = append(changes, Change{Type: Removed, Path: path1, Message: "Property removed"})
			continue
		}

		sep2 := getSep(ctx2)
		path2 := ctx2 + sep2 + name

		changes = append(changes, compareSchema(path1, path2, prop1, prop2)...)
	}
	for name := range s2.Properties {
		if _, ok := s1.Properties[name]; !ok {
			sep2 := getSep(ctx2)
			path2 := ctx2 + sep2 + name
			changes = append(changes, Change{Type: Added, Path: path2, Message: "Property added"})
		}
	}

	// Compare Items (for arrays)
	if s1.Items != nil && s2.Items != nil {
		changes = append(changes, compareSchema(ctx1+"[]", ctx2+"[]", s1.Items, s2.Items)...)
	}
	return changes
}

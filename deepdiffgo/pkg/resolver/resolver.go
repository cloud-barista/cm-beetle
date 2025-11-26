package resolver

import (
	"fmt"
	"strings"

	"github.com/cloud-barista/cm-beetle/deepdiffgo/pkg/model"
)

type Resolver struct {
	Spec *model.SwaggerSpec
}

func New(spec *model.SwaggerSpec) *Resolver {
	return &Resolver{Spec: spec}
}

// Resolve expands all $ref in the spec.
func (r *Resolver) Resolve() error {
	for _, item := range r.Spec.Paths {
		if err := r.resolveOperation(item.Get); err != nil {
			return err
		}
		if err := r.resolveOperation(item.Put); err != nil {
			return err
		}
		if err := r.resolveOperation(item.Post); err != nil {
			return err
		}
		if err := r.resolveOperation(item.Delete); err != nil {
			return err
		}
		if err := r.resolveOperation(item.Options); err != nil {
			return err
		}
		if err := r.resolveOperation(item.Head); err != nil {
			return err
		}
		if err := r.resolveOperation(item.Patch); err != nil {
			return err
		}

		for _, param := range item.Parameters {
			if err := r.resolveParameter(param); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Resolver) resolveOperation(op *model.Operation) error {
	if op == nil {
		return nil
	}
	for _, param := range op.Parameters {
		if err := r.resolveParameter(param); err != nil {
			return err
		}
	}

	if op.RequestBody != nil {
		for _, mediaType := range op.RequestBody.Content {
			if mediaType.Schema != nil {
				if err := r.resolveSchema(mediaType.Schema, map[string]bool{}); err != nil {
					return err
				}
			}
		}
	}

	for code, resp := range op.Responses {
		// resp is a value in the map, so we need to take address, modify, and put back?
		// No, range over map gives copy of value.
		// We need to modify the map.
		respCopy := resp
		if err := r.resolveResponse(&respCopy); err != nil {
			return err
		}
		op.Responses[code] = respCopy
	}
	return nil
}

func (r *Resolver) resolveParameter(param *model.Parameter) error {
	if param == nil {
		return nil
	}
	if param.Ref != "" {
		resolved, err := r.lookupParameter(param.Ref)
		if err != nil {
			return err
		}
		*param = *resolved
		param.Ref = ""
	}

	if param.Schema != nil {
		if err := r.resolveSchema(param.Schema, map[string]bool{}); err != nil {
			return err
		}
	}
	if param.Items != nil {
		if err := r.resolveSchema(param.Items, map[string]bool{}); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resolver) resolveResponse(resp *model.Response) error {
	if resp.Ref != "" {
		resolved, err := r.lookupResponse(resp.Ref)
		if err != nil {
			return err
		}
		*resp = *resolved
		resp.Ref = ""
	}
	if resp.Schema != nil {
		if err := r.resolveSchema(resp.Schema, map[string]bool{}); err != nil {
			return err
		}
	}
	for _, mediaType := range resp.Content {
		if mediaType.Schema != nil {
			if err := r.resolveSchema(mediaType.Schema, map[string]bool{}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Resolver) resolveSchema(schema *model.Schema, visited map[string]bool) error {
	if schema == nil {
		return nil
	}

	if schema.Ref != "" {
		if visited[schema.Ref] {
			// Cycle detected, stop recursion
			return nil
		}
		visited[schema.Ref] = true

		resolved, err := r.lookupSchema(schema.Ref)
		if err != nil {
			return err
		}

		// Deep copy resolved schema to avoid modifying the original definition
		// and to allow independent resolution of nested refs
		schemaCopy := r.deepCopySchema(resolved)

		// Preserve the original ref name for reporting
		parts := strings.Split(schema.Ref, "/")
		if len(parts) > 0 {
			schemaCopy.RefName = parts[len(parts)-1]
		}

		// Overwrite current schema with resolved one
		*schema = *schemaCopy
		schema.Ref = ""

		// Continue resolving recursively
		if err := r.resolveSchema(schema, visited); err != nil {
			return err
		}

		// We don't remove from visited here because in this branch of recursion, this ref is active.
	}

	for _, prop := range schema.Properties {
		if err := r.resolveSchema(prop, copyMap(visited)); err != nil {
			return err
		}
	}

	if schema.Items != nil {
		if err := r.resolveSchema(schema.Items, copyMap(visited)); err != nil {
			return err
		}
	}

	// Handle AllOf, OneOf, AnyOf if needed
	for _, s := range schema.AllOf {
		if err := r.resolveSchema(s, copyMap(visited)); err != nil {
			return err
		}
	}

	return nil
}

func (r *Resolver) lookupSchema(ref string) (*model.Schema, error) {
	parts := strings.Split(ref, "/")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid ref: %s", ref)
	}

	if parts[1] == "definitions" {
		if s, ok := r.Spec.Definitions[parts[2]]; ok {
			return &s, nil
		}
	} else if parts[1] == "components" && parts[2] == "schemas" {
		if r.Spec.Components != nil {
			if s, ok := r.Spec.Components.Schemas[parts[3]]; ok {
				return &s, nil
			}
		}
	}
	return nil, fmt.Errorf("schema ref not found: %s", ref)
}

func (r *Resolver) lookupParameter(ref string) (*model.Parameter, error) {
	parts := strings.Split(ref, "/")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid ref: %s", ref)
	}
	if parts[1] == "parameters" {
		if p, ok := r.Spec.Parameters[parts[2]]; ok {
			return &p, nil
		}
	} else if parts[1] == "components" && parts[2] == "parameters" {
		if r.Spec.Components != nil {
			if p, ok := r.Spec.Components.Parameters[parts[3]]; ok {
				return &p, nil
			}
		}
	}
	return nil, fmt.Errorf("parameter ref not found: %s", ref)
}

func (r *Resolver) lookupResponse(ref string) (*model.Response, error) {
	parts := strings.Split(ref, "/")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid ref: %s", ref)
	}
	if parts[1] == "responses" {
		if resp, ok := r.Spec.Responses[parts[2]]; ok {
			return &resp, nil
		}
	} else if parts[1] == "components" && parts[2] == "responses" {
		if r.Spec.Components != nil {
			if resp, ok := r.Spec.Components.Responses[parts[3]]; ok {
				return &resp, nil
			}
		}
	}
	return nil, fmt.Errorf("response ref not found: %s", ref)
}

func (r *Resolver) deepCopySchema(s *model.Schema) *model.Schema {
	if s == nil {
		return nil
	}
	out := *s

	if s.Properties != nil {
		out.Properties = make(map[string]*model.Schema)
		for k, v := range s.Properties {
			out.Properties[k] = r.deepCopySchema(v)
		}
	}
	if s.Items != nil {
		out.Items = r.deepCopySchema(s.Items)
	}
	if s.AllOf != nil {
		out.AllOf = make([]*model.Schema, len(s.AllOf))
		for i, v := range s.AllOf {
			out.AllOf[i] = r.deepCopySchema(v)
		}
	}
	// ... handle OneOf, AnyOf ...
	return &out
}

func copyMap(m map[string]bool) map[string]bool {
	n := make(map[string]bool)
	for k, v := range m {
		n[k] = v
	}
	return n
}

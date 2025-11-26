package model

// SwaggerSpec represents the root of the Swagger 2.0 / OpenAPI 3.0 document
type SwaggerSpec struct {
	Swagger     string               `json:"swagger,omitempty" yaml:"swagger,omitempty"`
	OpenAPI     string               `json:"openapi,omitempty" yaml:"openapi,omitempty"`
	Info        Info                 `json:"info" yaml:"info"`
	Host        string               `json:"host,omitempty" yaml:"host,omitempty"`
	BasePath    string               `json:"basePath,omitempty" yaml:"basePath,omitempty"`
	Schemes     []string             `json:"schemes,omitempty" yaml:"schemes,omitempty"`
	Consumes    []string             `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces    []string             `json:"produces,omitempty" yaml:"produces,omitempty"`
	Paths       map[string]PathItem  `json:"paths" yaml:"paths"`
	Definitions map[string]Schema    `json:"definitions,omitempty" yaml:"definitions,omitempty"` // Swagger 2.0
	Components  *Components          `json:"components,omitempty" yaml:"components,omitempty"`   // OpenAPI 3.0
	Parameters  map[string]Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`   // Global parameters
	Responses   map[string]Response  `json:"responses,omitempty" yaml:"responses,omitempty"`     // Global responses
}

type Info struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Version     string `json:"version" yaml:"version"`
}

type Components struct {
	Schemas    map[string]Schema    `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Parameters map[string]Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses  map[string]Response  `json:"responses,omitempty" yaml:"responses,omitempty"`
}

// PathItem represents operations available on a single path
type PathItem struct {
	Get        *Operation   `json:"get,omitempty" yaml:"get,omitempty"`
	Put        *Operation   `json:"put,omitempty" yaml:"put,omitempty"`
	Post       *Operation   `json:"post,omitempty" yaml:"post,omitempty"`
	Delete     *Operation   `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options    *Operation   `json:"options,omitempty" yaml:"options,omitempty"`
	Head       *Operation   `json:"head,omitempty" yaml:"head,omitempty"`
	Patch      *Operation   `json:"patch,omitempty" yaml:"patch,omitempty"`
	Parameters []*Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Ref        string       `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type Operation struct {
	Summary     string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	OperationID string                `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Consumes    []string              `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces    []string              `json:"produces,omitempty" yaml:"produces,omitempty"`
	Parameters  []*Parameter          `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody *RequestBody          `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses   map[string]Response   `json:"responses" yaml:"responses"`
	Schemes     []string              `json:"schemes,omitempty" yaml:"schemes,omitempty"`
	Deprecated  bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security    []map[string][]string `json:"security,omitempty" yaml:"security,omitempty"`
	Tags        []string              `json:"tags,omitempty" yaml:"tags,omitempty"`
}

type RequestBody struct {
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Required    bool                 `json:"required,omitempty" yaml:"required,omitempty"`
	Ref         string               `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type Parameter struct {
	Name             string      `json:"name,omitempty" yaml:"name,omitempty"`
	In               string      `json:"in,omitempty" yaml:"in,omitempty"`
	Description      string      `json:"description,omitempty" yaml:"description,omitempty"`
	Required         bool        `json:"required,omitempty" yaml:"required,omitempty"`
	Type             string      `json:"type,omitempty" yaml:"type,omitempty"`
	Format           string      `json:"format,omitempty" yaml:"format,omitempty"`
	Items            *Schema     `json:"items,omitempty" yaml:"items,omitempty"`
	CollectionFormat string      `json:"collectionFormat,omitempty" yaml:"collectionFormat,omitempty"`
	Default          interface{} `json:"default,omitempty" yaml:"default,omitempty"`
	Schema           *Schema     `json:"schema,omitempty" yaml:"schema,omitempty"` // Body parameter
	Ref              string      `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type Response struct {
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Schema      *Schema              `json:"schema,omitempty" yaml:"schema,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Ref         string               `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

type MediaType struct {
	Schema *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

type Schema struct {
	Ref                  string             `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Type                 string             `json:"type,omitempty" yaml:"type,omitempty"`
	Format               string             `json:"format,omitempty" yaml:"format,omitempty"`
	Description          string             `json:"description,omitempty" yaml:"description,omitempty"`
	Default              interface{}        `json:"default,omitempty" yaml:"default,omitempty"`
	Enum                 []interface{}      `json:"enum,omitempty" yaml:"enum,omitempty"`
	Properties           map[string]*Schema `json:"properties,omitempty" yaml:"properties,omitempty"`
	Items                *Schema            `json:"items,omitempty" yaml:"items,omitempty"`
	Required             []string           `json:"required,omitempty" yaml:"required,omitempty"`
	AdditionalProperties interface{}        `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"` // bool or Schema
	Example              interface{}        `json:"example,omitempty" yaml:"example,omitempty"`
	AllOf                []*Schema          `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf                []*Schema          `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf                []*Schema          `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	RefName              string             `json:"-" yaml:"-"` // Internal use: original ref name
}

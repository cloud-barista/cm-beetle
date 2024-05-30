package model

type Response struct {
	Success bool                   `json:"success" example:"true"`
	Text    string                 `json:"text" example:"Any text"`
	Detail  string                 `json:"details,omitempty" example:"Any details"`
	Object  map[string]interface{} `json:"object,omitempty"`
	List    []interface{}          `json:"list,omitempty"`
}

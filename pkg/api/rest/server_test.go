package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAllowedOrigins(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want []string
	}{
		{name: "Wildcard", raw: "*", want: []string{"*"}},
		{name: "Single origin", raw: "http://localhost:3000", want: []string{"http://localhost:3000"}},
		{name: "Comma-separated with spaces", raw: "http://localhost:3000, https://beetle.example.com , http://127.0.0.1:8080", want: []string{"http://localhost:3000", "https://beetle.example.com", "http://127.0.0.1:8080"}},
		{name: "Trailing and duplicate commas", raw: ",http://a.example,,http://b.example,", want: []string{"http://a.example", "http://b.example"}},
		{name: "Empty and whitespace", raw: " ,  , ", want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, parseAllowedOrigins(tt.raw))
		})
	}
}

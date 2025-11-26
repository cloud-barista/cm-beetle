package loader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloud-barista/cm-beetle/deepdiffgo/pkg/model"
	"gopkg.in/yaml.v3"
)

// Load reads a Swagger/OpenAPI file from the given path or URL and returns the parsed spec.
func Load(path string) (*model.SwaggerSpec, error) {
	var data []byte
	var err error

	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resp, err := http.Get(path)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch URL %s: %w", path, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch URL %s: status code %d", path, resp.StatusCode)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body from %s: %w", path, err)
		}
	} else {
		data, err = os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", path, err)
		}
	}

	var spec model.SwaggerSpec
	ext := strings.ToLower(filepath.Ext(path))

	if ext == ".json" {
		if err := json.Unmarshal(data, &spec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
	} else if ext == ".yaml" || ext == ".yml" {
		if err := yaml.Unmarshal(data, &spec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
		}
	} else {
		// Try YAML first, then JSON
		if err := yaml.Unmarshal(data, &spec); err != nil {
			if err := json.Unmarshal(data, &spec); err != nil {
				return nil, fmt.Errorf("failed to unmarshal as YAML or JSON: %w", err)
			}
		}
	}

	return &spec, nil
}

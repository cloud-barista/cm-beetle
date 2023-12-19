package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	Init()
}

func Init() {
	viper.AddConfigPath("../../conf/") // config for development
	viper.AddConfigPath(".")           // config for production optionally looking for the configuration in the working directory
	viper.AddConfigPath("./conf/")     // config for production optionally looking for the configuration in the working directory/conf/
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Map environment variable names to config file key names
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Automatically recognize environment variables
	viper.AutomaticEnv()

	// Try to read the config file in development environment
	if viper.GetString("node.env") != "production" {
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading config file, using default settings: %s", err)
		}
	}

	// Values set in runtime
	if viper.GetString("cmbeetle.root") == "" {
		log.Println("cmbeetle.root is not set in config file or environment variable")

		log.Println("find project root by using project name")
		projectName := "cm-beetle"
		projectRoot, err := findProjectRoot(projectName)
		if err != nil {
			log.Fatalf("Error finding project root directory: %v", err)
		}
		// Set the binary path
		viper.Set("cmbeetle.root", projectRoot)
		viper.Set("cbstore.root", projectRoot)
		viper.Set("cblog.root", projectRoot)
		viper.Set("apidoc.path", projectRoot+"/pkg/api/rest/docs/swagger.json")
	}

	// Recursively print all keys and values in Viper
	settings := viper.AllSettings()
	recursivePrintMap(settings, "")
}

func findProjectRoot(projectName string) (string, error) {
	// Get the executable path
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error getting executable path: %v", err)
	}
	execDir := filepath.Dir(execPath)

	// find last index of project name
	index := strings.LastIndex(execDir, projectName)
	if index == -1 {
		log.Println("project name not found the path")
		return "", errors.New("proejct name not found in the path")
	}

	// Cut the string up to the index
	result := execDir[:index+len(projectName)]

	log.Printf("project root directory: %s\n", result)

	return result, nil
}

func recursivePrintMap(m map[string]interface{}, prefix string) {
	for k, v := range m {
		fullKey := prefix + k
		if nestedMap, ok := v.(map[string]interface{}); ok {
			// Recursive call for nested maps
			recursivePrintMap(nestedMap, fullKey+".")
		} else {
			// Print current key-value pair
			log.Printf("Key: %s, Value: %v\n", fullKey, v)
		}
	}
}

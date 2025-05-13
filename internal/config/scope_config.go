package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// APIResourceConfig holds the configuration for a single API resource,
// including its path, HTTP method, and the scopes required to access it.
type APIResourceConfig struct {
	Path   string   `yaml:"path"`
	Method string   `yaml:"method"`
	Scopes []string `yaml:"scopes"`
}

// ScopeConfig holds a list of APIResourceConfig, representing the
// scope configurations for all protected API resources.
type ScopeConfig struct {
	APIResources []APIResourceConfig `yaml:"api_resources"`
}

var loadedScopeConfig *ScopeConfig

// LoadScopeConfig reads the API scope configuration from the specified YAML file path,
// unmarshals it into the ScopeConfig struct, and stores it for application use.
// It returns an error if reading or unmarshaling fails.
func LoadScopeConfig(filePath string) (*ScopeConfig, error) {
	if loadedScopeConfig != nil {
		return loadedScopeConfig, nil
	}

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config ScopeConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	loadedScopeConfig = &config
	return loadedScopeConfig, nil
}

// GetScopeConfig returns the loaded API scope configuration.
// It's recommended to call LoadScopeConfig once at application startup.
func GetScopeConfig() *ScopeConfig {
	return loadedScopeConfig
}

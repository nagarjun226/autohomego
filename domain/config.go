package domain

import (
	"encoding/json"
	"fmt"
)

// Config - Struct that will hold the config object
// Wrapper around the config
type Config struct {
	config map[string]interface{}
}

// SetFromJSONParsed - Set Config from read JSON file
func (c *Config) SetFromJSONParsed(b []byte) error {
	var rawConfig map[string]interface{}
	if err := json.Unmarshal(b, &rawConfig); err != nil {
		return err
	}

	c.config = rawConfig
	return nil
}

// GetConfig - obtain the config of a service. Returns the entire config if no argument provided
func (c *Config) GetConfig(service string) (map[string]interface{}, error) {
	if service == "" {
		return c.config, nil
	}

	baseConfig, ok := c.config["base"]
	if !ok {
		return nil, fmt.Errorf("Base config not present")
	}

	var baseConfigM map[string]interface{}
	baseConfigM, ok = baseConfig.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Base Config is not a map")
	}

	config, ok := c.config[service]
	if !ok {
		return baseConfigM, fmt.Errorf("Service specific config not present. Base config returned")
	}

	var configM map[string]interface{}
	configM, ok = config.(map[string]interface{})
	if !ok {
		return baseConfigM, fmt.Errorf("Service Config is not a map. Base config returned")
	}

	cM := make(map[string]interface{})
	for k, v := range baseConfigM {
		cM[k] = v
	}
	for k, v := range configM {
		cM[k] = v
	}

	return cM, nil
}

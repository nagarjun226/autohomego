package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// Config - Struct that will hold the config object
// Wrapper around the config
type Config struct {
	config map[string]interface{} //
}

// ConfigAutoLoader - keep reloading the given config object from location
// -- Will be deprecated when we add in the feature to edit config via PUT request
type ConfigAutoLoader struct {
	Config   *Config // Pointer to a config object that we choose to auto reload
	Location string  // location of the config.json to read from
	Rr       int     // refresh rate of config in seconds
}

// ALErrorChan - A channel that will handle errors thrown by ConfigAutoLoader.Run()
type ALErrorChan chan error

// HandleError - handle errors coming in on the channel
// For now doing nothing. Maybe later can print to a file?
func (c *ALErrorChan) HandleError() {
	<-*c
}

// AutoLoaderErrorChan - instance of ALErrorChan buffer length = 2
var AutoLoaderErrorChan = make(ALErrorChan)

// Run - run the autoconfig loader that re-reads the config.json every few seconds
func (c *ConfigAutoLoader) Run() {
	var err error
	go AutoLoaderErrorChan.HandleError() // spin out a new goroutine to handle errors that are not fatal
	for {
		// Handle Not set errors
		if c.Location == "" {
			err = fmt.Errorf("No Config location set yet. time = %v", time.Now())
			AutoLoaderErrorChan <- err
			continue
		}
		if c.Config == nil {
			err = fmt.Errorf("No Config set yet. time = %v", time.Now())
			AutoLoaderErrorChan <- err
			continue
		}
		if c.Rr == 0 {
			err = fmt.Errorf("No Refresh Rate Set set yet. time = %v", time.Now())
			AutoLoaderErrorChan <- err
			continue
		}

		// Read config.json and set the config
		configBytes, er := ioutil.ReadFile(c.Location)
		if er != nil {
			err = fmt.Errorf("Error while Reading File. %v", er)
			AutoLoaderErrorChan <- err
			continue
		}
		c.Config.SetFromJSONParsed(configBytes)

		time.Sleep(time.Duration(c.Rr) * time.Second) // Sleep untill it is time to set the json again

	}
}

// SetFromJSONParsed - Set Config from read JSON file
func (c *Config) SetFromJSONParsed(b []byte) error {

	// Parse the json into a map[string]interface{}
	var rawConfig map[string]interface{}
	if err := json.Unmarshal(b, &rawConfig); err != nil {
		return err
	}

	c.config = rawConfig
	return nil
}

// GetConfig - obtain the config of a service. Returns the entire config if no argument provided
func (c *Config) GetConfig(service string) (map[string]interface{}, error) {
	// If no service name provided, return the whole config
	if service == "" {
		return c.config, nil
	}

	// get the base config
	baseConfig, ok := c.config["base"]
	if !ok {
		return nil, fmt.Errorf("Base config not present")
	}

	// assert and typecase the baseconfig to map[string]interface{}
	var baseConfigM map[string]interface{}
	baseConfigM, ok = baseConfig.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Base Config is not a map")
	}

	// Do the same for the device specific config
	config, ok := c.config[service]
	if !ok {
		return baseConfigM, fmt.Errorf("Service specific config not present. Base config returned")
	}

	var configM map[string]interface{}
	configM, ok = config.(map[string]interface{})
	if !ok {
		return baseConfigM, fmt.Errorf("Service Config for service %v is not a map. Base config returned", service)
	}

	// Merge the configs. Device config takes precision
	cM := make(map[string]interface{})
	for k, v := range baseConfigM {
		cM[k] = v
	}
	for k, v := range configM {
		cM[k] = v
	}

	return cM, nil
}

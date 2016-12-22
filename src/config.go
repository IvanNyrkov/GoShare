package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Config stores settings of application
type Config struct {
	Port string `json:"port"` // Port on which runs this app
}

// InitConfig creates new Config struct filled with values from config file
func (app *App) InitConfig(configFilePath string) error {
	// Check environment variable
	if configFilePath == "" {
		return errors.New("Environment variable CONFIG_FILE (path to config file) should be specified.")
	}
	// Read bytes from file
	configFileContent, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	// Parse bytes into structure
	if err := json.Unmarshal(configFileContent, &app.Config); err != nil {
		return err
	}
	return nil
}

package main

import "os"

// Config stores settings of application
type Config struct {
	ServerPort string // Port on which runs this app
}

// NewConfig creates new Config struct filled with values from environment variables
func NewConfig() *Config {
	return &Config{
		ServerPort: os.Getenv("EXPOSE_PORT"),
	}
}

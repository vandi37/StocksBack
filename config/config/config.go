package config

import (
	"os"

	"github.com/VandiKond/vanerrors"
	"gopkg.in/yaml.v3"
)

// The errors
const (
	ErrorReadingConfig      = "error reading config"
	ErrorUnmarshalingConfig = "error unmarshaling config"
)

// The database connection config
type DatabaseCfg struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// The standard config
type StandardCfg struct {
	Port     string      `yaml:"port"`
	Database DatabaseCfg `yaml:"database"`
	Salt     string      `yaml:"salt"`
}

// Loads config from the yaml file
func LoadConfig(path string) (*StandardCfg, error) {
	// Getting the config data
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, vanerrors.NewWrap(ErrorReadingConfig, err, vanerrors.EmptyHandler)

	}

	// unmarshal the config
	var config StandardCfg
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, vanerrors.NewWrap(ErrorUnmarshalingConfig, err, vanerrors.EmptyHandler)
	}

	return &config, nil
}
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type App struct {
	Name        string `yaml:"name"`
	Port        string `yaml:"port"`
	Environment string `yaml:"environment"`
	Host        string `yaml:"host"`
}

type Config struct {
	App App `yaml:"app"`
}

func LoadConfigs(path string) (*Config, error) {
	var config Config

	yamlData, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("error occured while reading the yaml config file")
	}

	err = yaml.Unmarshal(yamlData, &config)

	if err != nil {
		return nil, fmt.Errorf("error occured while parsing the yaml data")
	}

	return &config, nil
}

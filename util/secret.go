package util

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config will contain all the configurables for the berfdei bot
type Config struct {
	AuthToken     string `yaml:"authToken"`
	CommandPrefix string `yaml:"commandPrefix"`
}

// ReadTokenFromFile reads the auth token from a yaml file
func ReadTokenFromFile(fileName string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("yamlFile.Get err   #%v ", err.Error())
	}
	var c *Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %v", err.Error())
	}
	return c, nil
}

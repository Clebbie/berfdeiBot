package util

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AuthToken string `yaml:"authToken"`
}

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

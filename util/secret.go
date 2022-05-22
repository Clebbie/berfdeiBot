package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
	authToken string
}

func (c *config) ReadToken() {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type myData struct {
	Config struct {
		RequestType string `yaml:"requestType"`
	}
}

func GetConf(fileName string) (*myData, error) {
	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	c := &myData{}
	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", fileName, err)
	}

	return c, err
}

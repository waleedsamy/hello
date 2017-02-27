package config

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Port  int    `json:"port" yaml:"port"`
	Env   string `json:"env" yaml:"env"`
	Mongo struct {
		URL      string `json:"url" yaml:"url"`
		Database string `json:"database" yaml:"database"`
	} `json:"mongo" yaml:"mongo"`
}

func GetConfig(file string) (*Config, error) {
	filename, err := filepath.Abs(file)
	if err != nil {
		return &Config{}, nil
	}
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return &Config{}, err
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return &Config{}, err
	}

	return &config, nil
}

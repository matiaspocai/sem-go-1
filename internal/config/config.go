package config

// go run cmd/chat/chatsrv.go -config ./config/config.yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// DbConfig ...
type DbConfig struct {
	Type   string `yaml:"type"`
	Driver string `yaml:"driver"`
	Conn   string `yaml:"conn"`
}

// Config ...
type Config struct {
	DB      DbConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

// LoadConfig ...
func LoadConfig(filename string) *Config {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var c = &Config{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		panic(err)
	}

	return c
}

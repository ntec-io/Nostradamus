package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents the config.yml
type Config struct {
	APIKey string `yaml:"APIKey"`
}

// ReadConfig TODO
func ReadConfig(fn string) (cfg Config, err error) {
	file, err := ioutil.ReadFile(fn)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(file, &cfg)
	return
}

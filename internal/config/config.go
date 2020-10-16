package config

import (
	"io/ioutil"

	"github.com/ntec-io/Nostradamus/internal/logger"
	"gopkg.in/yaml.v2"
)

// Config represents the config.yml
type Config struct {
	APIKey        string `yaml:"api_key"`
	RedisPassword string `yaml:"redis_password"`
}

// ReadConfig TODO
func ReadConfig(fn string) (cfg Config, err error) {
	logger.Log().Debug("Start reading config file")
	file, err := ioutil.ReadFile(fn)
	if err != nil {
		logger.Log().Error(err)
		return
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		logger.Log().Error(err)
		return
	}
	logger.Log().Debug("Config successfull read")
	return
}

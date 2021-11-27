package configs

import (
	"cloud-native-exercise/http-server/internal/configs/apps"
	"cloud-native-exercise/http-server/internal/configs/logs"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	App apps.AppConfigModel `yaml:"app"`
	Log logs.LogConfigModel `yaml:"log"`
}

var config *Config

func GetConfig() *Config {
	return config
}

func LoadConfig(file string) error {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		errorInfo := fmt.Sprintf("ReadFile %s failed, error:%s", file, err)
		log.Print(errorInfo)
		return errors.New(errorInfo)
	}

	if config == nil {
		config = &Config{}
	}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		errorInfo := fmt.Sprintf("Parser Config %s failed, error:%s", file, err)
		log.Print(errorInfo)
		return errors.New(errorInfo)
	}

	err = config.Parser()
	if err != nil {
		errorInfo := fmt.Sprintf("Parser Config %s Parser failed, error:%s", file, err)
		log.Print(errorInfo)
		return errors.New(errorInfo)
	}

	return nil
}

func (cfg *Config) Parser() error {
	return nil
}

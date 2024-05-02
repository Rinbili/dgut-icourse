package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var Config = initConf()

type Yaml struct {
	Secret string `yaml:"secret"`
	Port   string `yaml:"port"`
	Sql    struct {
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		Name   string `yaml:"name"`
	}
	Wechat struct {
		AppID     string `yaml:"appID"`
		AppSecret string `yaml:"appSecret"`
	}
}

func initConf() *Yaml {
	config := new(Yaml)
	yamlFile, err := os.ReadFile("./config/config.yml")
	if err != nil {
		panic("config.yml not found")
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic("config.yml unmarshal fail")
	}
	return config
}

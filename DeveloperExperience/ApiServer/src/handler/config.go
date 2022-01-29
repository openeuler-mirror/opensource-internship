package handler

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	ConfigContent Config
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type Config struct {
	ElasticSearch ElasticSearch `yaml:"elasticsearch"`
}

type ElasticSearch struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (c *Config) GetConf() *Config {
	yamlFile, err := ioutil.ReadFile("config/config.yml")
	CheckError(err)
	err = yaml.Unmarshal(yamlFile, c)
	CheckError(err)
	return c
}

func InitConfig() {
	ConfigContent.GetConf()
}

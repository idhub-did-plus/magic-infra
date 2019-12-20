package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ConfigType struct {
	Contract struct {
		Provider    string `yaml:"provider"`
		Identity    string `yaml:"identity"`
		Address780  string `yaml:"address780"`
		Address1484 string `yaml:"address1484"`
		Issuer      string `yaml:"issuer"`
	}
	Mysql struct {
		Url string `yaml:"url"`
	}
}

var Config ConfigType

func init() {
	content, _ := ioutil.ReadFile("./config.yaml")
	Config = ConfigType{}
	err := yaml.Unmarshal(content, &Config)

	fmt.Println(err, Config)
	return
}

package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ConfigType struct {
	Provider        string `yaml:"provider"`
	Identity        string `yaml:"identity"`
	ContractAddress string `yaml:"contractAddress"`
	Issuer          string `yaml:"issuer"`
}

var Config ConfigType

func init() {
	content, _ := ioutil.ReadFile("../config.yaml")
	Config = ConfigType{}
	err := yaml.Unmarshal(content, &Config)

	fmt.Println(err, Config)
	return
}

package configuration

import (
	"fmt"

	"github.com/space-fold-technologies/aurora-agent/app/core/server"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Host             string `yaml:"host"`
	Port             int    `yaml:"port"`
	Provider         string `yaml:"provider"`
	ProfileDIR       string `yaml:"profile-directory"`
	ListenAddress    string `yaml:"listen-address"`
	AdvertiseAddress string `yaml:"advertise-address"`
}

func ParseFromResource() Configuration {
	yamlFile, err := server.Asset("resources/settings.yml")
	config := Configuration{}
	if err != nil {
		fmt.Println("Failed to open the file")
		panic(err)
	}
	yaml.Unmarshal(yamlFile, &config)
	return config
}

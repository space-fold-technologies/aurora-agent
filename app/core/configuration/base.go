package configuration

import (
	"fmt"
	"os"

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

func ParseFromPath(filePath string) Configuration {
	config := Configuration{}
	if data, err := os.ReadFile(filePath); err != nil {
		fmt.Println("Failed to open the file")
		panic(err)
	} else if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}
	return config
}

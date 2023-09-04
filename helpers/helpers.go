package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server    ServerConfig     `yaml:"server"`
	Endpoints []EndpointConfig `yaml:"endpoints"`
}

type ServerConfig struct {
	Port  int  `yaml:"port"`
	Debug bool `yaml:"debug"`
}

type EndpointConfig struct {
	Name     string `yaml:"name"`
	Endpoint string `yaml:"endpoint"`
	Rc       int    `yaml:"rc"`
}

func ReadConfigurationFromYaml() Config {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return Config{}
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return Config{}
	}

	return config
}

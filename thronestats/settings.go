package thronestats

import (
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)


type ServerSettings struct {
	ListenAddress string		`yaml:"listen_address"`
	ListenPort    int           `yaml:"listen_port"`
	WwwPath       string        `yaml:"www_path"`
}


func (s *ServerSettings) ToYaml() []byte {
	result, err := yaml.Marshal(s)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return result
}

func GetServerSettings() ServerSettings {
	settings := ServerSettings{
		"0.0.0.0",
		8080,
		"../../www/",
	}

	data, err := ioutil.ReadFile("settings.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	yaml.Unmarshal(data, &settings)

	return settings
}

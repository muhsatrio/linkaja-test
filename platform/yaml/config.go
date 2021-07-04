package yaml

import (
	"os"

	"gopkg.in/yaml.v2"
)

func Open(fileName string) (config Config, err error) {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return
	}

	return
}

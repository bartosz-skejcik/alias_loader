package load

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Aliases map[string]map[string]string

func readFile(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}
	return data, nil
}

func parseYAML(data []byte, out interface{}) error {
	if err := yaml.Unmarshal(data, out); err != nil {
		return fmt.Errorf("error parsing YAML: %v", err)
	}
	return nil
}

func LoadAliases(yamlFilePath string) (Aliases, error) {
	data, err := readFile(yamlFilePath)
	if err != nil {
		return nil, err
	}

	var aliases Aliases
	if err := parseYAML(data, &aliases); err != nil {
		return nil, err
	}
	return aliases, nil
}

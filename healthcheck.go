package healthcheck

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type HealthCheckItem struct {
	Name    string            `yaml:"name"`
	Url     string            `yaml:"url"`
	Method  string            `yaml:"method"`
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func ParseYaml(filename string) ([]HealthCheckItem, error) {
	// Create a struct to hold the YAML data
	var items []HealthCheckItem

	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return items, err
	}

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &items)
	if err != nil {
		fmt.Println(err)
		return items, err
	}

	return items, nil
}

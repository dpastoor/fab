package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Setting struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Prompt  string `yaml:"prompt"`
	Default string `yaml:"default"`
}
type Config struct {
	Settings []Setting `yaml:"settings"`
	// should be a single dir
	Templates []string `yaml:"templates"`
	// will instead look at all dir's within the collection dir
	// for any that have a _setup.yml
	Collections []string `yaml:"collections"`
}

func (s Setting) Validate() error {
	if s.Name == "" {
		return errors.New("must provide a `name`")
	}
	switch s.Type {
	case "string", "boolean", "bool", "":
	default:
		return fmt.Errorf("invalid setting type %s", s.Type)

	}
	return nil
}

func (cfg Config) Validate() error {
	for _, s := range cfg.Settings {
		err := s.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func Read(path string) (Config, error) {
	var config Config
	file, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}
	err = config.Validate()
	if err != nil {
		return config, err
	}
	return config, nil
}

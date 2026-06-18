package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Reminder struct {
	Message string `yaml:"message"`
	Time    string `yaml:"schedule"`
}

type Config struct {
	Reminders []Reminder `yaml:"reminders"`
}

func readFile(fileName string) (reminder []Reminder, err error) {

	var config Config
	rawFile, err := os.ReadFile(fileName)
	if err != nil {
		return []Reminder{}, err
	}

	err1 := yaml.Unmarshal(rawFile, &config)

	if err1 != nil {
		return []Reminder{}, err
	}

	return config.Reminders, nil
}

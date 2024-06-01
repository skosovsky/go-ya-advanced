package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

func (s *Settings) Save(filename string) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal settings: %w", err)
	}

	if err = os.WriteFile(filename, data, 0600); err != nil {
		return fmt.Errorf("could not save settings: %w", err)
	}

	return nil
}

func (s *Settings) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read settings: %w", err)
	}

	err = json.Unmarshal(data, s)
	if err != nil {
		return fmt.Errorf("could not unmarshal settings: %w", err)
	}

	return nil
}

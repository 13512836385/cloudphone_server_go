package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	AK string `json:"ak"`
	SK string `json:"sk"`
}

func loadAKAndSK() (string, string, error) {
	homeDir, err := os.Getwd()
	if err != nil {
		return "", "", err
	}
	configPath := filepath.Join(homeDir, ".volc", "config")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", "", err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return "", "", err
	}
	return config.AK, config.SK, nil
}

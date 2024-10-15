package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Secrets struct {
	TwelvedataApiKey string `json:"twelvedataApiKey"`
}

func loadSecrets(filename string) (*Secrets, error) {
	var secretsLocal Secrets

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s file: %w", filename, err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s file: %w", filename, err)
	}

	if err := json.Unmarshal(byteValue, &secretsLocal); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &secretsLocal, nil
}

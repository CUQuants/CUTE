package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var secrets Secrets

type Secrets struct {
	TwelvedataApiKey string `json:"twelvedataApiKey"`
}

func loadSecrets(filename string) (Secrets, error) {
	var secrets Secrets

	file, err := os.Open(filename)
	if err != nil {
		return secrets, fmt.Errorf("failed to open %s file: %w", filename, err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return secrets, fmt.Errorf("failed to read %s file: %w", filename, err)
	}

	if err := json.Unmarshal(byteValue, &secrets); err != nil {
		return secrets, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return secrets, nil
}

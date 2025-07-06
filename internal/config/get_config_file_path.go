package config

import (
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	cfgPath := filepath.Join(homeDir, configFileName)
	return cfgPath, nil
}

package config

import (
	"os"
	"path/filepath"
)

var cfg Config

type Config struct {
	ConfigDir       string
	CandidateDir    string
	JenvDir         string
	JenvVersionsDir string
}

func Default() Config {
	homeDir, _ := os.UserHomeDir()
	return Config{
		ConfigDir:       filepath.Join(homeDir, ".config", "jdk-go"),
		CandidateDir:    filepath.Join(homeDir, ".config", "jdk-go", "candidates"),
		JenvDir:         filepath.Join(homeDir, ".jenv"),
		JenvVersionsDir: filepath.Join(homeDir, ".jenv", "versions"),
	}
}

func Init(c Config) {
	cfg = c
	if _, err := os.Stat(cfg.ConfigDir); os.IsNotExist(err) {
		os.Mkdir(cfg.ConfigDir, os.ModePerm)
	}
	if _, err := os.Stat(cfg.CandidateDir); os.IsNotExist(err) {
		os.Mkdir(cfg.ConfigDir, os.ModePerm)
	}
}

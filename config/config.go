package config

import (
	"os"
	"path/filepath"
)

var cfg Config

type Config struct {
	Dir             string
	CandidatesDir   string
	JenvDir         string
	JenvVersionsDir string
}

func Dir() string {
	return cfg.Dir
}

func CandidatesDir() string {
	return cfg.CandidatesDir
}

func JenvDir() string {
	return cfg.JenvDir
}

func JenvVersionsDir() string {
	return cfg.JenvVersionsDir
}

func Default() Config {
	homeDir, _ := os.UserHomeDir()
	return Config{
		Dir:             filepath.Join(homeDir, ".config", "jdk-go"),
		CandidatesDir:   filepath.Join(homeDir, ".config", "jdk-go", "candidates"),
		JenvDir:         filepath.Join(homeDir, ".jenv"),
		JenvVersionsDir: filepath.Join(homeDir, ".jenv", "versions"),
	}
}

func Init(c Config) {
	cfg = c
	if _, err := os.Stat(cfg.Dir); os.IsNotExist(err) {
		os.Mkdir(cfg.Dir, os.ModePerm)
	}
	if _, err := os.Stat(cfg.CandidatesDir); os.IsNotExist(err) {
		os.Mkdir(cfg.Dir, os.ModePerm)
	}
}

package jdkutil

import (
	"github.com/TheCheerfulDev/jdk-go/config"
	"os"
	"path/filepath"
	"strings"
)

const versionFile string = ".java-version"

var cfg config.Config

func GetConfigDir() string {
	return cfg.ConfigDir
}

func GetCandidatesDir() string {
	return cfg.CandidateDir
}

func GetJenvVersionsDir() string {
	return cfg.JenvVersionsDir
}

func GetJenvDir() string {
	return cfg.JenvDir
}

func GetActiveVersion() (version, versionFilePath string, err error) {
	currentDirectory, err := os.Getwd()

	if err != nil {
		return "", "", err
	}

	for {
		if !strings.HasSuffix(currentDirectory, "/") {
			currentDirectory += "/"
		}

		versionFilePath = currentDirectory + versionFile
		if _, err := os.Stat(currentDirectory + versionFile); !os.IsNotExist(err) {
			return ExtractActiveVersionFromFile(versionFilePath), versionFilePath, nil
		}

		if currentDirectory == "/" {
			homeDir, _ := os.UserHomeDir()
			versionFilePath = homeDir + "/.jenv/version"
			return ExtractActiveVersionFromFile(versionFilePath), versionFilePath, nil
		}

		currentDirectory = filepath.Clean(filepath.Join(currentDirectory, ".."))
	}

}

func ExtractActiveVersionFromFile(filePath string) (version string) {
	fileContent, _ := os.ReadFile(filePath)
	version = string(fileContent)
	version = RemoveNewLineFromString(version)
	return version
}

func RemoveNewLineFromString(input string) string {
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	return input
}

func Init(c config.Config) {
	cfg = c
}

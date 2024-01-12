package jdkutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const versionFile string = ".java-version"

var configDir string

func GetConfigDir() string {
	if configDir == "" {
		homeDir, _ := os.UserHomeDir()
		configDir = homeDir + "/.config/jdk-go"
	}
	return configDir
}

func GetCandidatesDir() string {
	return GetConfigDir() + "/candidates"
}

func GetJenvVersionsDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.jenv/versions"
}

func GetActiveVersion() (version, versionFilePath string) {
	currentDirectory, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		if !strings.HasSuffix(currentDirectory, "/") {
			currentDirectory += "/"
		}

		versionFilePath = currentDirectory + versionFile
		if _, err := os.Stat(currentDirectory + versionFile); !os.IsNotExist(err) {
			return ExtractActiveVersionFromFile(versionFilePath), versionFilePath
		}

		if currentDirectory == "/" {
			homeDir, _ := os.UserHomeDir()
			versionFilePath = homeDir + "/.jenv/version"
			return ExtractActiveVersionFromFile(versionFilePath), versionFilePath
		}

		currentDirectory = filepath.Clean(filepath.Join(currentDirectory, ".."))
	}

	return "nope", versionFilePath

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

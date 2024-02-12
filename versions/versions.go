package versions

import (
	"errors"
	"fmt"
	"github.com/TheCheerfulDev/jdk/config"
	"github.com/TheCheerfulDev/jdk/jdkutil"
	"os"
	"path/filepath"
	"strings"
)

const versionFile = ".java-version"

func RemoveLocalVersion() (err error) {
	dir, _ := os.Getwd()
	err = os.Remove(dir + "/.java-version")

	if err != nil {
		return err
	}

	return nil
}

func SetOrShowGlobalVersion(args []string) (err error) {
	if len(args) == 0 {
		fileContent, err := os.ReadFile(config.JenvDir() + "/version")
		if err != nil {
			return errors.New("No global JDK version defined")
		}
		globalVersion := string(fileContent)
		globalVersion = jdkutil.RemoveNewLineFromString(globalVersion)
		fmt.Println(globalVersion)
		return nil
	}
	version := args[0]

	if _, err := os.Stat(config.Dir() + "/" + version); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("JDK version %v does not exist", version))
	}

	os.WriteFile(config.JenvDir()+"/version", []byte(version), os.ModePerm)
	return nil
}

func SetOrShowLocalVersion(args []string) (err error) {
	if len(args) == 0 {
		dir, _ := os.Getwd()
		if fileContent, err := os.ReadFile(dir + "/.java-version"); !os.IsNotExist(err) {
			fmt.Println(string(fileContent))
			return nil
		}
		return errors.New("No local JDK version defined in this directory")
	}

	version := args[0]

	if _, err := os.Stat(config.Dir() + "/" + version); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("JDK version %v does not exist", version))
	}

	os.WriteFile(".java-version", []byte(version), os.ModePerm)
	return nil
}

func Active() (version, versionFilePath string, err error) {
	currentDirectory, _ := os.Getwd()

	for {
		if !strings.HasSuffix(currentDirectory, "/") {
			currentDirectory += "/"
		}

		versionFilePath = currentDirectory + versionFile
		if _, err := os.Stat(versionFilePath); !os.IsNotExist(err) {
			versionInFile, err := extractActiveVersionFromFile(versionFilePath)
			return versionInFile, versionFilePath, err
		}

		if currentDirectory == "/" {
			versionFilePath = filepath.Join(config.JenvDir(), "version")
			versionInFile, err := extractActiveVersionFromFile(versionFilePath)
			return versionInFile, versionFilePath, err
		}

		currentDirectory = filepath.Clean(filepath.Join(currentDirectory, ".."))
	}
}

func Remove(version string) (err error, alias string, hasAlias bool) {
	if _, err := os.Stat(config.Dir() + "/" + version); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("JDK version %v does not exist", version)), "", false
	}

	hasAlias, aliasToRemove := aliasForVersion(version)

	if hasAlias {
		_ = os.Remove(config.JenvVersionsDir() + "/" + aliasToRemove)
		_ = os.Remove(config.Dir() + "/" + aliasToRemove)
	}

	_ = os.Remove(config.JenvVersionsDir() + "/" + version)
	_ = os.Remove(config.Dir() + "/" + version)
	_ = os.RemoveAll(config.CandidatesDir() + "/" + version)

	return nil, aliasToRemove, hasAlias
}

func aliasForVersion(version string) (bool, string) {
	configDir := config.Dir()

	files, _ := os.ReadDir(configDir)

	for _, file := range files {
		if !IsVersionFile(file) {
			continue
		}

		fileContent, _ := os.ReadFile(configDir + "/" + file.Name())
		versionInFile := string(fileContent)
		if versionInFile == version {
			return true, file.Name()
		}

	}

	return false, ""
}

func IsVersionFile(file os.DirEntry) bool {
	return !file.IsDir() && !strings.HasPrefix(file.Name(), ".")
}

func extractActiveVersionFromFile(filePath string) (version string, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	version = string(fileContent)
	version = jdkutil.RemoveNewLineFromString(version)
	return version, nil
}

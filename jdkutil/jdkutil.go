package jdkutil

import (
	"os/exec"
	"strings"
)

func GetVersionFromJenv() string {
	return strings.Split(GetFullVersionFromJenv(), " ")[0]
}

func GetFullVersionFromJenv() string {
	commandOutput, _ := exec.Command("jenv", "version").Output()
	return string(commandOutput)
}

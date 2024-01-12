package main

import (
	"github.com/TheCheerfulDev/jdk-go/cmd"
	"github.com/TheCheerfulDev/jdk-go/jdkutil"
	"os"
)

func main() {
	if _, err := os.Stat(jdkutil.GetConfigDir()); os.IsNotExist(err) {
		os.Mkdir(jdkutil.GetConfigDir(), os.ModePerm)
	}
	if _, err := os.Stat(jdkutil.GetCandidatesDir()); os.IsNotExist(err) {
		os.Mkdir(jdkutil.GetCandidatesDir(), os.ModePerm)
	}
	cmd.Execute()
}

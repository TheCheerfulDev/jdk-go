package main

import (
	"github.com/TheCheerfulDev/jdk-go/cmd"
	"os"
)

func main() {
	if _, err := os.Stat("/Users/mark/.config/jdk2"); os.IsNotExist(err) {
		os.Mkdir("/Users/mark/.config/jdk2", os.ModePerm)
	}
	cmd.Execute()
}

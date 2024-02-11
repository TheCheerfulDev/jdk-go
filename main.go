package main

import (
	"github.com/TheCheerfulDev/jdk-go/cmd"
	"github.com/TheCheerfulDev/jdk-go/config"
	"github.com/TheCheerfulDev/jdk-go/jdkutil"
)

func main() {
	c := config.Default()
	config.Init(c)
	jdkutil.Init(c)
	cmd.Execute()
}

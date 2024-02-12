package main

import (
	"github.com/TheCheerfulDev/jdk/cmd"
	"github.com/TheCheerfulDev/jdk/config"
)

func main() {
	c := config.Default()
	config.Init(c)
	cmd.Execute()
}

package cmd

import (
	"errors"
	"fmt"
	"github.com/TheCheerfulDev/jdk/config"
	"github.com/TheCheerfulDev/jdk/versions"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var configDir = config.Dir()
var activeVersion string

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all installed JDK versions",
	Long: `As the naming convention implies, this command lists all installed JDK versions.

If a version is an alias, it will be displayed after ->.
The currently active version will be preceded by an asterisk (*).`,
	Run: func(cmd *cobra.Command, args []string) {
		err := handleLs()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func handleLs() error {
	var err error
	activeVersion, _, err = versions.Active()
	if err != nil {
		return errors.New("Could not read the active version")
	}

	configDir = config.Dir()
	files, _ := os.ReadDir(configDir)

	fmt.Println("Installed JDKs:")

	for _, file := range files {
		if versions.IsVersionFile(file) {
			printVersionInformation(file)
		}
	}
	return nil
}

func printVersionInformation(file os.DirEntry) {
	fileInfo, _ := file.Info()
	readFile, _ := os.ReadFile(configDir + "/" + file.Name())
	prefixText := getPrefixText(file.Name())

	if fileInfo.Size() > 0 {
		fmt.Println(prefixText, file.Name(), "             ->", strings.ReplaceAll(string(readFile), "\n", ""))
		return
	}

	fmt.Println(prefixText, file.Name())
}

func getPrefixText(version string) interface{} {
	if version == activeVersion {
		return "*"
	}
	return " "
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

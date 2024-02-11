package cmd

import (
	"fmt"
	"github.com/TheCheerfulDev/jdk-go/jdkutil"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var configDir = jdkutil.GetConfigDir()
var activeVersion string

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all installed JDK versions",
	Long: `As the naming convention implies, this command lists all installed JDK versions.

If a version is an alias, it will be displayed after ->.
The currently active version will be preceded by an asterisk (*).`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		activeVersion, _, err = jdkutil.GetActiveVersion()
		if err != nil {
			fmt.Println("Could not read the active version")
			os.Exit(1)
		}

		configDir = jdkutil.GetConfigDir()
		files, err := os.ReadDir(configDir)
		if err != nil {
			fmt.Println("Could not read the config directory")
			os.Exit(1)
		}

		fmt.Println("Installed JDKs:")

		for _, file := range files {
			if IsVersionFile(file) {
				printVersionInformation(file)
			}
		}

	},
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

func IsVersionFile(file os.DirEntry) bool {
	return !file.IsDir() && !strings.HasPrefix(file.Name(), ".")
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

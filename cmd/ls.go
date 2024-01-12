package cmd

import (
	"fmt"
	"github.com/TheCheerfulDev/jdk-go/jdkutil"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var homeDir string
var configDir string
var activeVersion string

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all installed JDK versions.",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		activeVersion, _ = jdkutil.GetActiveVersion()

		// Get all the files in jdk2 folder
		homeDir, _ = os.UserHomeDir()
		configDir = homeDir + "/.config/jdk2"
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
	readFile, _ := os.ReadFile(homeDir + "/.config/jdk2/" + file.Name())
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

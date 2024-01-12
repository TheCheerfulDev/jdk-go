package cmd

import (
	"fmt"
	"github.com/TheCheerfulDev/jdk-go/jdkutil"
	"os"

	"github.com/spf13/cobra"
)

// localCmd represents the local command
var localCmd = &cobra.Command{
	Use:               "local [JDK version]",
	ValidArgsFunction: CustomVersionCompletion,
	Args:              cobra.MaximumNArgs(1),
	Short:             "Set or show the local application-specific JDK",
	Long: `This function sets the local application-specific JDK to the
provided version. You can also use an alias.

Example usage:
	jdk-go local 21`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			dir, _ := os.Getwd()
			if fileContent, err := os.ReadFile(dir + "/.java-version"); !os.IsNotExist(err) {
				fmt.Println(string(fileContent))
				os.Exit(0)
			}
			fmt.Println("No local JDK version defined in this directory")
			os.Exit(1)
		}

		version := args[0]

		if _, err := os.Stat(jdkutil.GetConfigDir() + "/" + version); os.IsNotExist(err) {
			fmt.Printf("JDK version %v does not exist\n", version)
			os.Exit(1)
		}

		os.WriteFile(".java-version", []byte(version), os.ModePerm)
	},
}

func init() {
	rootCmd.AddCommand(localCmd)
}

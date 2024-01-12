package cmd

import (
	"fmt"
	"github.com/TheCheerfulDev/jdk-go/jdkutil"
	"os"

	"github.com/spf13/cobra"
)

// globalCmd represents the global command
var globalCmd = &cobra.Command{
	Use:               "global [JDK version]",
	ValidArgsFunction: CustomVersionCompletion,
	Args:              cobra.MaximumNArgs(1),
	Short:             "Set or show the global JDK",
	Long: `This function sets the global JDK to the
provided version. You can also use an alias.

Example usage:
	jdk-go global 21`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fileContent, _ := os.ReadFile(jdkutil.GetJenvDir() + "/version")
			fmt.Println(string(fileContent))
			os.Exit(0)
		}
		version := args[0]

		if _, err := os.Stat(jdkutil.GetConfigDir() + "/" + version); os.IsNotExist(err) {
			fmt.Printf("JDK version %v does not exist\n", version)
			os.Exit(1)
		}

		os.WriteFile(jdkutil.GetJenvDir()+"/version", []byte(version), os.ModePerm)
	},
}

func init() {
	rootCmd.AddCommand(globalCmd)
}

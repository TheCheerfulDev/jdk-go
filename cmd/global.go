package cmd

import (
	"fmt"
	"github.com/TheCheerfulDev/jdk/versions"
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

		err := versions.SetOrShowGlobalVersion(args)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(globalCmd)
}

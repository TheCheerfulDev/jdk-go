package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// localCmd represents the local command
var localCmd = &cobra.Command{
	Use:               "local",
	ValidArgsFunction: CustomVersionCompletion,
	Args:              cobra.ExactArgs(1),
	Short:             "Set the local application-specific JDK",
	Long: `This function sets the local application-specific JDK to the
provided version. You can also use an alias.

Example usage:
	jdk-go local 21`,
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		os.WriteFile(".java-version", []byte(version), os.ModePerm)
	},
}

func init() {
	rootCmd.AddCommand(localCmd)
}

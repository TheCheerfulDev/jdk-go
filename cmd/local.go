package cmd

import (
	"fmt"
	"github.com/TheCheerfulDev/jdk/versions"
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
	jdk local 21`,
	Run: func(cmd *cobra.Command, args []string) {
		err := versions.SetOrShowLocalVersion(args)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(localCmd)
}

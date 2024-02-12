package cmd

import (
	"fmt"
	"github.com/TheCheerfulDev/jdk/versions"
	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the application-specific JDK, if one is set in the current directory.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := versions.RemoveLocalVersion()

		if err != nil {
			fmt.Println("No application-specific JDK configuration found")
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}

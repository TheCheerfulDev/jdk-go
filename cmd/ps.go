package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "Show the currently active JDK and its jenv origin.",
	Long:  "This is the long stuff",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Active JDK:")

	},
}

func init() {
	rootCmd.AddCommand(psCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "jdk-go",
	Short:   "",
	Version: "0.1.0",
	Long: `This is a go implementation of the JDK wrapper.
This CLI app is used to manage your Java JDK installations.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(fmt.Sprintf("jdk version %s - © Mark Hendriks <thecheerfuldev>\n", rootCmd.Version))
}

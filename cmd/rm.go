package cmd

import (
	"fmt"
	"github.com/TheCheerfulDev/jdk/config"
	"github.com/TheCheerfulDev/jdk/versions"
	"github.com/spf13/cobra"
	"os"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Args:  cobra.ExactArgs(1),
	Short: "Remove a JDK version",
	Long: `This function removes the provided JDK, along with any alias it might have.

Example usage:
	jdk-go rm 21-tem`,
	ValidArgsFunction: CustomVersionCompletion,
	Run: func(cmd *cobra.Command, args []string) {
		versionToRemove := args[0]

		err, aliasToRemove, hasAlias := versions.Remove(versionToRemove)
		if err != nil {
			fmt.Println(err)
			return
		}

		printRemovalSuccesMessage(versionToRemove, aliasToRemove, hasAlias)
	},
}

func printRemovalSuccesMessage(versionToRemove string, aliasToRemove string, hasAlias bool) {
	if hasAlias {
		fmt.Printf("Succesfully removed JDK version %v and alias %v\n", versionToRemove, aliasToRemove)
		return
	}
	fmt.Printf("Succesfully removed JDK version %v\n", versionToRemove)
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

func CustomVersionCompletion(_ *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {

	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	dir, _ := os.ReadDir(config.Dir())

	versionList := make([]string, 2)

	for _, file := range dir {
		if versions.IsVersionFile(file) {
			versionList = append(versionList, file.Name())
		}
	}

	return versionList, cobra.ShellCompDirectiveNoFileComp
}

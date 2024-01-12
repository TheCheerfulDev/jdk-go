package cmd

import (
	"fmt"
	"github.com/TheCheerfulDev/jdk-go/jdkutil"
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
	Run: func(cmd *cobra.Command, args []string) {
		// check if jdk exists

		versionToRemove := args[0]

		if _, err := os.Stat(jdkutil.GetConfigDir() + "/" + versionToRemove); os.IsNotExist(err) {
			fmt.Printf("JDK %v does not exist\n", versionToRemove)
			return
		}

		getFilesToRemoveForVersion(versionToRemove)
		hasAlias, aliasToRemove := getAliasForVersion(versionToRemove)
		fmt.Println(hasAlias, aliasToRemove)

		// remove aliasToRemove from jenv
		if hasAlias {
			os.Remove(jdkutil.GetJenvVersionsDir() + "/" + aliasToRemove)
			// remove aliasToRemove file
			os.Remove(jdkutil.GetConfigDir() + "/" + aliasToRemove)
		}

		// remove version from jenv
		os.Remove(jdkutil.GetJenvVersionsDir() + "/" + versionToRemove)
		// remove version file
		os.Remove(jdkutil.GetConfigDir() + "/" + versionToRemove)

		// remove candidate version directory
		os.RemoveAll(jdkutil.GetCandidatesDir() + "/" + versionToRemove)

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

func getAliasForVersion(version string) (bool, string) {
	configDir := jdkutil.GetConfigDir()

	files, _ := os.ReadDir(configDir)

	for _, file := range files {
		if !IsVersionFile(file) {
			continue
		}

		fileContent, _ := os.ReadFile(configDir + "/" + file.Name())
		versionInFile := string(fileContent)
		if versionInFile == version {
			return true, file.Name()
		}

	}

	return false, ""
}

func getFilesToRemoveForVersion(remove string) {

}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

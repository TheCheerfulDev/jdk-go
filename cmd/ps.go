package cmd

import (
	"errors"
	"fmt"
	"github.com/TheCheerfulDev/jdk/config"
	"github.com/TheCheerfulDev/jdk/jdkutil"
	"github.com/TheCheerfulDev/jdk/versions"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Args:  cobra.NoArgs,
	Short: "Show the currently active JDK and its jenv origin.",
	Long:  "This is the long stuff",
	Run: func(cmd *cobra.Command, args []string) {
		err := handlePs()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func handlePs() error {
	activeVersion, path, err := versions.Active()
	if err != nil {
		return errors.New("Could not read the active version")
	}

	configDir = config.Dir()

	file, err := os.ReadFile(filepath.Join(configDir, activeVersion))

	if err != nil {
		return errors.New(fmt.Sprintf("Active JDK version %v does not exist", activeVersion))
	}

	fmt.Println("Active JDK:")
	fileContent := string(file)
	fileContent = jdkutil.RemoveNewLineFromString(fileContent)

	if fileContent != "" {
		fmt.Printf("  %v (set by %v) -> %v\n", activeVersion, path, fileContent)
		return nil
	}

	fmt.Printf("  %v (set by %v)\n", activeVersion, path)
	return nil
}

func init() {
	rootCmd.AddCommand(psCmd)
}

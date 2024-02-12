package cmd

import (
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
		handlePs()
	},
}

func handlePs() {
	activeVersion, path, err := versions.Active()

	configDir = config.Dir()

	file, err := os.ReadFile(filepath.Join(configDir, activeVersion))

	if err != nil {
		fmt.Printf("Active JDK version %v does not exist\n", activeVersion)
		os.Exit(1)
	}

	fmt.Println("Active JDK:")
	fileContent := string(file)
	fileContent = jdkutil.RemoveNewLineFromString(fileContent)

	if fileContent != "" {
		fmt.Printf("  %v (set by %v) -> %v\n", activeVersion, path, fileContent)
		return
	}

	fmt.Printf("  %v (set by %v)\n", activeVersion, path)
}

func init() {
	rootCmd.AddCommand(psCmd)
}

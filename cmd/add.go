package cmd

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/TheCheerfulDev/jdk/config"
	"github.com/TheCheerfulDev/jdk/jdkutil"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add [url of JDK tarball]",
	Aliases: []string{"install"},
	Short:   "Add a JDK from the provided URL",
	Args:    cobra.ExactArgs(1),
	Long: `This command adds a new JDK. For this to work you must provide an URL
pointing to a tarball (.tar.gz) of the JDK you want to add.

example:
	jdk add https://www.myjdk.com/21/jdk-21.tar.gz`,
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		if !strings.HasSuffix(url, ".tar.gz") {
			fmt.Println("The provided URL does not seem to be valid!")
			os.Exit(1)
		}

		var version string
		var alias string

		fmt.Print("Please provide the version of the JDK to add: ")
		fmt.Scanln(&version)
		version = jdkutil.RemoveNewLineFromString(version)

		if version == "" {
			fmt.Println("JDK version can't be empty")
			os.Exit(1)
		}

		fmt.Printf("Please provide alias for JDK version %v (leave empty for none): ", version)
		fmt.Scanln(&alias)
		alias = jdkutil.RemoveNewLineFromString(alias)

		if doesFileAlreadyExist(version) {
			fmt.Printf("JDK version %v already exists\n", version)
			os.Exit(1)
		}

		if doesFileAlreadyExist(alias) {
			fmt.Printf("JDK alias %v already exists\n", alias)
			os.Exit(1)
		}

		downloadJdkFromUrl(url, version)
		addJdk(version, alias)
		printSuccessMessage(version, alias)
	},
}

func doesFileAlreadyExist(fileName string) bool {
	if fileName == "" {
		return false
	}

	_, err := os.Stat(filepath.Join(config.Dir(), fileName))
	return !os.IsNotExist(err)
}

func printSuccessMessage(version, alias string) {

	if alias == "" {
		fmt.Printf("Successfully installed JDK version %v\n", version)
		return
	}
	fmt.Printf("Successfully installed JDK version %v with alias %v\n", version, alias)
}

func addJdk(version, alias string) {
	fileName := filepath.Join(config.CandidatesDir(), version) + ".tar.gz"
	destination := filepath.Join(config.CandidatesDir(), version)

	unTarJdk(fileName, destination)
	addVersion(version)
	addAlias(version, alias)
	createSimLinks(version)
	addJdkToJenv(version, alias)

}

func addJdkToJenv(version, alias string) {
	symlink := filepath.Join(config.JenvVersionsDir(), version)
	target := filepath.Join(config.CandidatesDir(), version)
	err := os.Symlink(target, symlink)

	if err != nil {
		fmt.Println(err)
		return
	}

	if alias != "" {
		symlink := config.JenvVersionsDir() + "/" + alias
		target := config.CandidatesDir() + "/" + version
		err := os.Symlink(target, symlink)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func createSimLinks(version string) {

	directories := []string{"conf", "include", "jmods", "legal", "lib", "bin", "man"}

	for _, directory := range directories {
		root := filepath.Join(config.CandidatesDir(), version)
		var target = ""

		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if d.Name() == directory {
				target = path
			}
			return nil
		})

		symlink := filepath.Join(config.CandidatesDir(), version, directory)
		if target != "" {
			os.Symlink(target, symlink)
		}

	}
}

func addVersion(version string) {
	versionFile, _ := os.Create(filepath.Join(config.Dir(), version))
	versionFile.Close()
}

func addAlias(version, alias string) {
	if alias == "" {
		return
	}
	_ = os.WriteFile(filepath.Join(config.Dir(), alias), []byte(version), 0644)
}

func unTarJdk(fileName, destination string) {
	open, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		_ = os.Remove(fileName)
		os.Exit(1)
	}

	gzipReader, err := gzip.NewReader(open)
	if err != nil {
		fmt.Println("Could not read JDK tarball")
		_ = os.Remove(fileName)
		os.Exit(1)
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()

		switch {
		case err == io.EOF:
			_ = os.Remove(fileName)
			return
		case err != nil:
			fmt.Println("Something went wrong while unpacking tarball")
			_ = os.Remove(fileName)
			_ = os.RemoveAll(destination)
			os.Exit(1)
		case header == nil:
			continue
		}
		target := filepath.Join(destination, header.Name)

		switch header.Typeflag {

		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					fmt.Println("Something went wrong while unpacking tarball")
					_ = os.Remove(fileName)
					_ = os.RemoveAll(destination)
					os.Exit(1)
				}
			}

		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				fmt.Println("Something went wrong while unpacking tarball")
				_ = os.Remove(fileName)
				_ = os.RemoveAll(destination)
				os.Exit(1)
			}

			if _, err := io.Copy(f, tarReader); err != nil {
				fmt.Println("Something went wrong while unpacking tarball")
				_ = os.Remove(fileName)
				_ = os.RemoveAll(destination)
				os.Exit(1)
			}

			f.Close()
		}

	}
}

func downloadJdkFromUrl(url, version string) {
	fileName := config.CandidatesDir() + "/" + version + ".tar.gz"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Downloading failed!")
		os.Exit(1)
	}
	defer resp.Body.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"Downloading JDK",
	)
	f, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	_, err = io.Copy(io.MultiWriter(f, bar), resp.Body)

	if err != nil {
		fmt.Println("Downloading failed!")
		_ = os.Remove(fileName)
		os.Exit(1)
	}

	defer f.Close()

}

func init() {
	addCmd.SetUsageTemplate(`
Usage: 
  jdk-go add [url of JDK tarball]

Aliases:
  add, install

Resources for downloading JDKs:
  Amazon Corretto: https://aws.amazon.com/corretto/
  Azul Zulu:       https://www.azul.com/downloads/
  OpenJDK:         https://jdk.java.net/
  Oracle JDK:      https://www.oracle.com/java/technologies/downloads/
  Temurin:         https://adoptium.net/temurin/releases/
`)
	rootCmd.AddCommand(addCmd)
}

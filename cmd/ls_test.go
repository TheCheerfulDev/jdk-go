package cmd

import (
	"github.com/TheCheerfulDev/jdk/config"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestHandleLs(t *testing.T) {
	_ = initLsTest(t)

	err := handleLs()
	assert.ErrorContains(t, err, "Could not read the active version")
}

func TestHandleLsWithInstalledJDKs(t *testing.T) {
	_ = initLsTest(t)

	tempDir := t.TempDir()
	_ = os.Chdir(tempDir)

	os.WriteFile(filepath.Join(tempDir, ".java-version"), []byte("21"), 0644)

	os.WriteFile(filepath.Join(config.Dir(), "21-tem"), []byte(""), 0644)
	os.WriteFile(filepath.Join(config.Dir(), "21"), []byte("21-tem"), 0644)

	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	handleLs()

	_ = w.Close()
	os.Stdout = stdout

	result, _ := io.ReadAll(r)
	output := string(result)
	assert.Contains(t, output, "Installed JDKs:")
	assert.Contains(t, output, "* 21")
	assert.Contains(t, output, "21-tem")
	assert.Contains(t, output, "-> 21-tem")
}

func initLsTest(t *testing.T) config.Config {
	tempDir := t.TempDir()

	c := config.Config{
		Dir:             filepath.Join(tempDir, ".config", "jdk-go"),
		CandidatesDir:   filepath.Join(tempDir, ".config", "jdk-go", "candidates"),
		JenvDir:         filepath.Join(tempDir, ".jenv"),
		JenvVersionsDir: filepath.Join(tempDir, ".jenv", "versions"),
	}
	config.Init(c)
	return c
}

package cmd

import (
	"github.com/TheCheerfulDev/jdk/config"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestHandlePs(t *testing.T) {
	_ = initPsTest(t)

	err := handlePs()
	assert.ErrorContains(t, err, "Could not read the active version")
}

func TestHandlePsWithExistingJDKAndAlias(t *testing.T) {
	_ = initPsTest(t)

	tempDir := t.TempDir()
	_ = os.Chdir(tempDir)

	os.WriteFile(filepath.Join(tempDir, ".java-version"), []byte("21"), 0644)

	os.WriteFile(filepath.Join(config.Dir(), "21-tem"), []byte(""), 0644)
	os.WriteFile(filepath.Join(config.Dir(), "21"), []byte("21-tem"), 0644)

	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := handlePs()

	_ = w.Close()
	os.Stdout = stdout

	result, _ := io.ReadAll(r)
	output := string(result)
	assert.NoError(t, err)
	assert.Contains(t, output, "Active JDK:")
	assert.Contains(t, output, "21 (set by")
	assert.Contains(t, output, "-> 21-tem")
}

func TestHandlePsWithExistingJDK(t *testing.T) {
	_ = initPsTest(t)

	tempDir := t.TempDir()
	_ = os.Chdir(tempDir)

	os.WriteFile(filepath.Join(tempDir, ".java-version"), []byte("21-tem"), 0644)

	os.WriteFile(filepath.Join(config.Dir(), "21-tem"), []byte(""), 0644)

	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := handlePs()

	_ = w.Close()
	os.Stdout = stdout

	result, _ := io.ReadAll(r)
	output := string(result)
	assert.NoError(t, err)
	assert.Contains(t, output, "Active JDK:")
	assert.Contains(t, output, "21-tem (set by")
	assert.NotContains(t, output, "->")
}

func TestHandlePsWithNonExistingJDK(t *testing.T) {
	_ = initPsTest(t)

	tempDir := t.TempDir()
	_ = os.Chdir(tempDir)

	os.WriteFile(filepath.Join(tempDir, ".java-version"), []byte("21-tem"), 0644)

	err := handlePs()

	assert.ErrorContains(t, err, "Active JDK version 21-tem does not exist")
}

func initPsTest(t *testing.T) config.Config {
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

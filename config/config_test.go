package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestDefault(t *testing.T) {
	homeDir, _ := os.UserHomeDir()
	cfg := Default()
	assert.Equal(t, filepath.Join(homeDir, ".config", "jdk-go"), cfg.Dir)
	assert.Equal(t, filepath.Join(homeDir, ".config", "jdk-go", "candidates"), cfg.CandidatesDir)
	assert.Equal(t, filepath.Join(homeDir, ".jenv"), cfg.JenvDir)
	assert.Equal(t, filepath.Join(homeDir, ".jenv", "versions"), cfg.JenvVersionsDir)
}

func TestInit(t *testing.T) {
	c := initConfigTest(t)
	_, err := os.Stat(c.Dir)
	assert.NoError(t, err)
	_, err = os.Stat(c.CandidatesDir)
	assert.NoError(t, err)
}

func TestDir(t *testing.T) {
	c := initConfigTest(t)
	assert.Equal(t, c.Dir, Dir())
}

func TestCandidatesDir(t *testing.T) {
	c := initConfigTest(t)
	assert.Equal(t, c.CandidatesDir, CandidatesDir())
}

func TestJenvDir(t *testing.T) {
	c := initConfigTest(t)
	assert.Equal(t, c.JenvDir, JenvDir())
}

func TestJenvVersionsDir(t *testing.T) {
	c := initConfigTest(t)
	assert.Equal(t, c.JenvVersionsDir, JenvVersionsDir())
}

func initConfigTest(t *testing.T) Config {
	tempDir := t.TempDir()

	c := Config{
		Dir:             filepath.Join(tempDir, ".config", "jdk-go"),
		CandidatesDir:   filepath.Join(tempDir, ".config", "jdk-go", "candidates"),
		JenvDir:         filepath.Join(tempDir, ".jenv"),
		JenvVersionsDir: filepath.Join(tempDir, ".jenv", "versions"),
	}
	Init(c)
	return c
}

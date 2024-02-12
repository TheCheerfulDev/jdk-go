package versions

import (
	"github.com/TheCheerfulDev/jdk/config"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveLocalVersion(t *testing.T) {
	_ = InitVersionsTest(t)
	tempDir := t.TempDir()
	_ = os.Chdir(tempDir)

	// Test no version
	err := RemoveLocalVersion()
	assert.Error(t, err)

	// Test local version exists
	_, _ = os.Create(".java-version")
	err = RemoveLocalVersion()
	assert.NoError(t, err)

	assert.NoFileExists(t, ".java-version")
}

func TestSetOrShowLocalVersion(t *testing.T) {
	_ = InitVersionsTest(t)
	tempDir := t.TempDir()
	_ = os.Chdir(tempDir)

	// Test no version
	err := SetOrShowLocalVersion([]string{})
	assert.ErrorContains(t, err, "No local JDK version defined in this directory")

	// Test set version
	_, _ = os.Create(filepath.Join(config.Dir(), "21"))
	err = SetOrShowLocalVersion([]string{"21"})
	assert.NoError(t, err)

	// Test show version
	err = SetOrShowLocalVersion([]string{})
	assert.NoError(t, err)

	// Test version does not exist
	err = SetOrShowLocalVersion([]string{"22"})
	assert.ErrorContains(t, err, "JDK version 22 does not exist")

}

func TestSetOrShowGlobalVersion(t *testing.T) {
	_ = InitVersionsTest(t)

	// Test no version
	err := SetOrShowGlobalVersion([]string{})
	assert.ErrorContains(t, err, "No global JDK version defined")

	// Test set version
	_, _ = os.Create(filepath.Join(config.Dir(), "21"))
	err = SetOrShowGlobalVersion([]string{"21"})
	assert.NoError(t, err)

	// Test show version
	err = SetOrShowGlobalVersion([]string{})
	assert.NoError(t, err)

	// Test version does not exist
	err = SetOrShowGlobalVersion([]string{"22"})
	assert.ErrorContains(t, err, "JDK version 22 does not exist")
}

func TestActive(t *testing.T) {
	_ = InitVersionsTest(t)
	tempDir := t.TempDir()
	_ = os.Chdir(tempDir)

	// Test no version
	_, _, err := Active()
	assert.ErrorContains(t, err, ".jenv/version: no such file or directory")

	// Test global active
	expectedVersion := "21"
	_ = os.WriteFile(filepath.Join(config.JenvDir(), "version"), []byte(expectedVersion), os.ModePerm)
	version, _, err := Active()
	assert.NoError(t, err)
	assert.Equal(t, expectedVersion, version)

	// Test local active
	expectedVersion = "22"
	_ = os.WriteFile(".java-version", []byte(expectedVersion), os.ModePerm)
	version, _, err = Active()
	assert.NoError(t, err)
	assert.Equal(t, expectedVersion, version)
}

func TestIsVersionFile(t *testing.T) {
	_ = InitVersionsTest(t)
	tempDir := t.TempDir()
	_ = os.Chdir(tempDir)
	_, _ = os.Create(".not-a-version-file")
	os.Mkdir("this-is-a-directory", os.ModePerm)
	dir, _ := os.ReadDir(tempDir)
	for _, entry := range dir {
		assert.False(t, IsVersionFile(entry))
	}

	_ = os.Remove(".not-a-version-file")
	_ = os.Remove("this-is-a-directory")

	_, _ = os.Create("is-a-version-file")
	dir, _ = os.ReadDir(tempDir)
	for _, entry := range dir {
		assert.True(t, IsVersionFile(entry))
	}
}

func TestAliasForVersion(t *testing.T) {
	_ = InitVersionsTest(t)
	tempDir := t.TempDir()
	_ = os.Chdir(tempDir)

	// Test Alias does not exist
	exists, _ := aliasForVersion("21")
	assert.False(t, exists)

	// Test Alias exists
	_ = os.WriteFile(filepath.Join(config.Dir(), "21"), []byte("21-tem"), os.ModePerm)
	exists, alias := aliasForVersion("21-tem")
	assert.True(t, exists)
	assert.Equal(t, "21", alias)
}

func TestRemove(t *testing.T) {
	_ = InitVersionsTest(t)
	tempDir := t.TempDir()
	_ = os.Chdir(tempDir)

	versionToRemove := "21-tem"
	alias := "21"

	// Test version does not exist
	err, _, _ := Remove(versionToRemove)
	assert.ErrorContains(t, err, "JDK version 21-tem does not exist")

	// Test version exists
	_ = os.WriteFile(filepath.Join(config.Dir(), alias), []byte(versionToRemove), os.ModePerm)
	_ = os.WriteFile(filepath.Join(config.Dir(), versionToRemove), []byte(""), os.ModePerm)
	err, removedAlias, hasAlias := Remove(versionToRemove)
	assert.NoError(t, err)
	assert.NoFileExists(t, filepath.Join(config.Dir(), versionToRemove))
	assert.NoFileExists(t, filepath.Join(config.Dir(), alias))
	assert.Equal(t, alias, removedAlias)
	assert.True(t, hasAlias)

}

func InitVersionsTest(t *testing.T) config.Config {
	tempDir := t.TempDir()

	c := config.Config{
		Dir:             filepath.Join(tempDir, ".config", "jdk-go"),
		CandidatesDir:   filepath.Join(tempDir, ".config", "jdk-go", "candidates"),
		JenvDir:         filepath.Join(tempDir, ".jenv"),
		JenvVersionsDir: filepath.Join(tempDir, ".jenv", "versions"),
	}
	config.Init(c)
	_ = os.MkdirAll(c.JenvVersionsDir, os.ModePerm)
	return c
}

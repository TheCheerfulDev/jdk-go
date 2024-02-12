package cmd

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestPrintRemovalSuccessMessageWithAlias(t *testing.T) {

	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printRemovalSuccesMessage("21-tem", "21", true)

	_ = w.Close()
	os.Stdout = stdout
	result, _ := io.ReadAll(r)
	output := string(result)
	assert.Contains(t, output, "Succesfully removed JDK version 21-tem and alias 21")

}

func TestPrintRemovalSuccessMessageNoAlias(t *testing.T) {

	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printRemovalSuccesMessage("21-tem", "", false)

	_ = w.Close()
	os.Stdout = stdout
	result, _ := io.ReadAll(r)
	output := string(result)
	assert.Contains(t, output, "Succesfully removed JDK version 21-tem")
	assert.NotContains(t, output, "alias")

}

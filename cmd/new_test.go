package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteEmptyCommand(t *testing.T) {
	var expectedOut bytes.Buffer
	var output bytes.Buffer

	newCmd.SetErr(&expectedOut)
	newCmd.Help()

	newCmd.SetErr(&output)
	newCmd.RunE(newCmd, nil)

	assert.Equal(t, expectedOut, output)
}

func TestMakeDirFunc(t *testing.T) {
	makeDir("foo")
	assert.DirExists(t, "foo")
	os.Remove("foo")
}

func TestMakeEmptyFileFunc(t *testing.T) {
	makeEmptyFile("foo")
	assert.FileExists(t, "foo")
	os.Remove("foo")
}

func TestInitDirs(t *testing.T) {
	initDirs("foo")
	for _, dir := range initialDirs {
		assert.DirExists(t, "foo"+PathSeparatorString+dir)
		assert.FileExists(t, "foo"+PathSeparatorString+dir+PathSeparatorString+".gitkeep")
	}
	os.RemoveAll("foo")
}

func TestInitialize(t *testing.T) {
	initialize(newCmd, []string{"foo"})
	assert.NoFileExists(t, "foo"+PathSeparatorString+".gitkeep")
	for _, dir := range initialDirs {
		assert.DirExists(t, "foo"+PathSeparatorString+dir)
		assert.FileExists(t, "foo"+PathSeparatorString+dir+PathSeparatorString+".gitkeep")
	}
	os.RemoveAll("foo")
}

func TestExecuteCommand(t *testing.T) {
	var expectedOut bytes.Buffer
	var output bytes.Buffer
	var args = []string{"foo"}

	newCmd.SetErr(&output)
	newCmd.RunE(newCmd, args)

	assert.Equal(t, expectedOut, output)
	os.RemoveAll("foo")
}

func TestExecuteCommandTooManyArgs(t *testing.T) {
	var expectedOut bytes.Buffer
	var output bytes.Buffer
	var args = []string{"foo", "bar"}

	newCmd.SetErr(&expectedOut)
	newCmd.Help()

	newCmd.SetErr(&output)
	newCmd.RunE(newCmd, args)

	assert.Equal(t, expectedOut, output)
}

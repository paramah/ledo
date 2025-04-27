package compose

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os/exec"
	"testing"

	"github.com/paramah/ledo/app/modules/context"
	"github.com/stretchr/testify/assert"
)

var (
	execCommand    = exec.Command
	cmdOutput      string
	ioutilReadFile = func(filename string) ([]byte, error) {
		return nil, errors.New("not implemented")
	}
)

func mockExecCommand(command string, args ...string) *exec.Cmd {
	return &exec.Cmd{
		Path:   command,
		Args:   append([]string{command}, args...),
		Stdout: &bytes.Buffer{},
	}
}

func captureOutput(f func()) string {
	var buf bytes.Buffer

	f()
	return buf.String()
}

func TestCheckDockerComposeVersion(t *testing.T) {
	execCommand = mockExecCommand
	defer func() { execCommand = exec.Command }()

	tests := []struct {
		name       string
		output     string
		shouldFail bool
	}{
		{"ValidVersion", "docker-compose version 1.29.2, build 5becea4c", false},
		{"InvalidVersion", "docker-compose version 1.27.0, build 5becea4c", true},
		{"NoVersion", "docker-compose", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdOutput = tt.output
			if tt.shouldFail {
				assert.Panics(t, CheckDockerComposeVersion)
			} else {
				assert.NotPanics(t, CheckDockerComposeVersion)
			}
		})
	}
}

func TestMergeComposerFiles(t *testing.T) {
	fileContent := `
version: '3'
services:
  web:
    image: nginx
`
	ioutilReadFile = func(filename string) ([]byte, error) {
		return []byte(fileContent), nil
	}
	defer func() { ioutilReadFile = ioutil.ReadFile }()

	result, err := MergeComposerFiles("file1.yaml", "file2.yaml")
	assert.NoError(t, err)

	expectedContent := `version: "3"
services:
    web:
        image: nginx
`
	assert.Equal(t, expectedContent, result)
}

func TestPrintCurrentMode(t *testing.T) {
	ctx := &context.LedoContext{
		Mode: context.Mode{CurrentMode: "test"},
	}

	output := captureOutput(func() {
		PrintCurrentMode(ctx)
	})

	assert.Contains(t, output, "MODE")
	assert.Contains(t, output, "test")
}

func TestExecComposerUp(t *testing.T) {
	ctx := &context.LedoContext{
		ComposeArgs: []string{},
		GetRuntimeCompose: func() string {
			return "docker-compose"
		},
		ExecCmd: func(cmd string, args ...string) error {
			return nil
		},
	}

	output := captureOutput(func() {
		ExecComposerUp(ctx, false)
	})

	assert.Contains(t, output, "up")
}

func TestExecComposerDown(t *testing.T) {
	ctx := &context.LedoContext{
		ComposeArgs: []string{},
		GetRuntimeCompose: func() string {
			return "docker-compose"
		},
		ExecCmd: func(cmd string, args ...string) error {
			return nil
		},
	}

	output := captureOutput(func() {
		ExecComposerDown(ctx)
	})

	assert.Contains(t, output, "down")
}

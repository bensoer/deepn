package utils

import (
	"bytes"
	"os/exec"
)

type CommandResult struct {
	ExitCode int
	Stdout   []byte
	Stderr   []byte
}

type Subproc struct {
}

func NewSubproc() *Subproc {
	return &Subproc{}
}

func (s *Subproc) Run(name string, args ...string) CommandResult {
	cmd := exec.Command(name, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Stdin = bytes.NewReader(nil)

	err := cmd.Run()

	result := CommandResult{
		Stdout: stdout.Bytes(),
		Stderr: stderr.Bytes(),
	}

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
		} else {
			// Process failed to start (e.g. command not found)
			result.ExitCode = -1
		}
	} else {
		result.ExitCode = 0
	}

	return result
}

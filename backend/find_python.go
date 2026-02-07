package main

import (
	"fmt"
	"os/exec"
)

func find_python() (string, error) {
	// 1. Try pyenv
	if _, err := exec.LookPath("pyenv"); err == nil {
		cmd := exec.Command("pyenv", "which", "python")
		out, err := cmd.Output()
		if err == nil {
			pythonPath := string(out)
			// remove any trailing newline
			if len(pythonPath) > 0 && pythonPath[len(pythonPath)-1] == '\n' {
				pythonPath = pythonPath[:len(pythonPath)-1]
			}
			return pythonPath, nil
		}

	}


	// 2. Try python3
	if path, err := exec.LookPath("python3"); err == nil {
		return path, nil
	}

	// 3. Try python
	if path, err := exec.LookPath("python"); err == nil {
		return path, nil
	}

	return "", fmt.Errorf("no Python executable found")
}

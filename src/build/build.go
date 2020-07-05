// Package build ...
package build

import (
	"fmt"
	"os/exec"
)

func Runfile() (string, error) {

	output, flag, input := "echo", "-n", "You successfully ran a Linux command from Go!!!"

	// the exec.Command makes sys calls to the Linux terminal
	// output returns the stdout from the terminal
	cmdLine, err := exec.Command(output, flag, input).Output()

	// test for error
	if err != nil {
		return "", fmt.Errorf("failed to run %q: %w", output, err)
	}

	// return converted output value
	return string(cmdLine), nil
}

func Native() (string, error) {
	output, err := exec.Command("/bin/sh", "/Users/richard/Desktop/Go/jekyll-server-build/src/build/echo.sh").Output()
	if err != nil {
		return "", fmt.Errorf("failed to run %q: %w", output, err)
	}
	return string(output), nil
}

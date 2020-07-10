package thing

import (
	"fmt"
	"os/exec"
	"strings"
)

func Command(command string, args ...string) (string, error) {
	// the exec.Command makes sys calls to the Linux terminal
	// output returns the stdout from the terminal
	output, err := exec.Command(command, args...).CombinedOutput()

	// test for error
	if err != nil {
		return "", fmt.Errorf("failed to run '%s %s': %w", command, strings.Join(args, " "), err)
	}

	// return converted output value
	return string(output), nil
}

func Shell(script string) (string, error) {
	output, err := exec.Command("/bin/sh", script).Output()
	if err != nil {
		return "", fmt.Errorf("failed to run %q: %w", script, err)
	}
	return string(output), nil
}

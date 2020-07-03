// Package build ...
package build

import (
	"fmt"
	"os/exec"
)

// Runfile function to execute Linux command
func Runfile() string {

	// command and input string
	cmd, input := "echo", "You successfully ran a Linux command from Go!!!"

	// the exec.Command makes sys calls to the Linux terminal
	// output returns the stdout from the terminal
	out, err := exec.Command(cmd, input).Output()

	// test for error
	if err != nil {
		fmt.Printf("%s", err)
	}
	// convert stdout to string
	output := string([]byte(out))

	// return converted output value
	return output
}

// Native function to call Linux command
func Native() int {

	return 0
}

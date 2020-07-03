// Package build ...
package build

import (
	"fmt"
	"os/exec"
)

// Runfile function to run Linux command
func Runfile() string {

	// command and input
	cmd, input := "echo", "You successfully ran a Linux command from Go!!!"

	// the exec.Command will make sys calls to the Linux terminal
	// Output returns the stdout form the terminal
	out, err := exec.Command(cmd, input).Output()

	// test for error
	if err != nil {
		fmt.Printf("%s", err)
	}
	// convert stdout to string
	//output := string(out[:])
	output := string([]byte(out))
	fmt.Print(output)

	// return converted output value
	return output
}

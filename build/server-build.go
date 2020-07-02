// Package main //server-build
package main

import (
	"fmt"
	"os/exec"
)

// Main function to run Linux command
func main() {

	// the exec.Command will make sys calls to the Linux terminal
	// Output returns the stdout form the terminal 
	out, err := exec.Command("echo", "You successfully ran a Linux command from Go!!!").Output()
	
	// test for error 
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	// convert stdout to string 
	output := string(out[:])
	fmt.Print(output)
	
	//cmd := exec.Command("echo", "You successfully ran a Linux command from Go!!!")
	//log.Printf("Running command")
	//err := cmd.Run()
	//log.Printf("Command finished with error: %v", err)
}
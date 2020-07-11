package thing

import (
	"fmt"
	"os/exec"
	"strings"
)

type Runner struct {
	Test    bool
	CmdLine string
}

func NewRunner() *Runner {
	return &Runner{Test: false}
}

func (r *Runner) Command(command string, args ...string) (string, error) {
	if r.Test {
		r.CmdLine = fmt.Sprintf("%s %s", command, strings.Join(args, " "))
		return "", nil
	}
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

func (r *Runner) YumUpdate() {
	r.Command("yum", "update -y")
}

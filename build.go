package thing

import (
	"fmt"
	"os/exec"
	"strings"
)

type Runner struct {
	Test    bool
	History []string
	Output  string
}

func NewRunner() *Runner {
	return &Runner{Test: false}
}

func (r *Runner) Command(command string, args ...string) error {
	r.History = append(r.History, fmt.Sprintf("%s %s", command, strings.Join(args, " ")))
	if r.Test {
		return nil
	}
	output, err := exec.Command(command, args...).CombinedOutput()
	r.Output = string(output)
	if err != nil {
		return fmt.Errorf("failed to run '%s %s': %w", command, strings.Join(args, " "), err)
	}
	return nil
}

func (r *Runner) UpdateYum() error {
	return r.Command("yum", "update -y")
}

func (r *Runner) InstallPackages(packages []string) error {
	for _, p := range packages {
		err := r.Command("yum", "install -y "+p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Runner) InstallGems(packages []string) error {
	for _, p := range packages {
		err := r.Command("gem", "install "+p)
		if err != nil {
			return err
		}
	}
	return nil
}

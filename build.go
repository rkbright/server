package thing

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Runner struct {
	test      bool
	History   []string
	Output    string
	updated   bool
	installed bool
}

func NewRunner() *Runner {
	return &Runner{}
}

func TestNewRunner() *Runner {
	r := NewRunner()
	r.test = true
	return r
}

func (r *Runner) Command(command string, args ...string) error {
	r.History = append(r.History, fmt.Sprintf("%s %s", command, strings.Join(args, " ")))
	if r.test {
		return nil
	}
	output, err := exec.Command(command, args...).CombinedOutput()
	r.Output = string(output)
	if err != nil {
		return fmt.Errorf("failed to run '%s %s': %w", command, strings.Join(args, " "), err)
	}
	return nil
}

func (r *Runner) InstallPackage(packages []string) error {

	if !r.updated {
		r.Command("yum", "update -y")
		r.updated = true
	}

	for _, p := range packages {
		if p != "update" {
			err := r.Command("yum", "install -y "+p)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Runner) InstallGems(packages []string) error {

	if !r.installed {
		log.Println("Need to install Ruby")
		r.InstallPackage([]string{"ruby"})
		r.installed = true

	}

	for _, p := range packages {
		err := r.Command("gem", "install "+p)
		if err != nil {
			return err
		}
	}
	return nil
}

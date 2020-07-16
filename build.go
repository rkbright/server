package thing

import (
	"fmt"
	"os/exec"
	"strings"
)

type Runner struct {
	Test    bool
	CmdLine string
	Output  string
}

func NewRunner() *Runner {
	return &Runner{Test: false}
}

func (r *Runner) Command(command string, args ...string) error {

	r.CmdLine = fmt.Sprintf("%s %s", command, strings.Join(args, " "))

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

func (r *Runner) YumUpdate() error {
	return r.Command("yum", "update -y")
}

func (r *Runner) YumInstall() error {
	return r.Command("yum", "install epel-release -y")
}

func (r *Runner) RubyInstall() error {
	return r.Command("yum", "install ruby -y")

}

func (r *Runner) JekyllInstall() error {
	return r.Command("gem", "install jekyll")
}

func (r *Runner) BundlerInstall() error {
	return r.Command("gem", "install bundler")
}

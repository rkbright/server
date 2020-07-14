package thing

import (
	"fmt"
	"os/exec"
	"strings"
)

type Runner struct {
	Test    bool
	CmdLine string
	output  string
}

func NewRunner() *Runner {
	return &Runner{Test: false}
}

func (r *Runner) Command(command string, args ...string) (string, error) {
	if r.Test {
		r.CmdLine = fmt.Sprintf("%s %s", command, strings.Join(args, " "))
		return "", nil
	}

	output, err := exec.Command(command, args...).CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("failed to run '%s %s': %w", command, strings.Join(args, " "), err)
	}

	return string(output), nil
}

func (r *Runner) YumUpdate() {
	r.Command("yum", "update -y")
	fmt.Printf("the output string: %s \n", r.CmdLine)
}

func (r *Runner) YumInstall() {
	r.Command("yum", "install epel-release -y")
}

func (r *Runner) RubyInstall() {
	r.Command("yum", "install ruby -y")
}

func (r *Runner) JekyllInstall() {
	r.Command("gem", "install jekyll")
}

func (r *Runner) BundlerInstall() {
	r.Command("gem", "install bundler")
}

func (r *Runner) RubyVersion() {
	r.Command("ruby", "2.6")
}

func (r *Runner) JekyllVersion() {
	r.Command("jekyll", "4.*")
}

func (r *Runner) BundlerVersion() {
	r.Command("Bundler", "2.*")
}

package server

import (
	"fmt"
	"os/exec"
	"strings"
)

type Runner struct {
	History       []string
	Output        string
	dryRun        bool
	yumUpdated    bool
	rubyInstalled bool
	installed     bool
}

func NewRunner() *Runner {
	return &Runner{}
}

func NewTestRunner() *Runner {
	r := NewRunner()
	r.dryRun = true
	return r
}

func (r *Runner) Command(command string, args ...string) error {
	r.History = append(r.History, fmt.Sprintf("%s %s", command, strings.Join(args, " ")))
	if r.dryRun {
		return nil
	}
	output, err := exec.Command(command, args...).CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		return fmt.Errorf("failed to run '%s %s': %w", command, strings.Join(args, " "), err)
	}
	return nil
}

func (r *Runner) InstallPackage(p string) error {
	err := r.EnsureYumUpdated()
	if err != nil {
		return err
	}
	err = r.Command("sudo", "yum", "install", "-y", p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) EnsureYumUpdated() error {
	if !r.yumUpdated {
		err := r.Command("sudo", "yum", "update", "-y")
		if err != nil {
			return err
		}
		r.yumUpdated = true
	}
	return nil
}

func (r *Runner) InstallGem(p string) error {
	err := r.EnsureRubyInstalled()
	if err != nil {
		return err
	}
	err = r.Command("gem", "install", p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) EnsureRubyInstalled() error {
	if !r.rubyInstalled {
		err := r.InstallPackage("epel-release")
		if err != nil {
			return err
		}
		err = r.InstallPackage("ruby")
		if err != nil {
			return err
		}
		r.rubyInstalled = true
	}
	return nil
}

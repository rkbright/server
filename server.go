package server

import (
	"fmt"
	"os/exec"
	"strings"
)

const rvmDependencies string = "httpd certbot python2-certbot-apache curl git-core gcc-c++ patch readline readline-devel zlib zlib-devel libffi-devel openssl-devel make bzip2 autoconf automake libtool bison sqlite-devel"

type Runner struct {
	History      []string
	Output       string
	Error        error
	dryRun       bool
	yumUpdated   bool
	rvmInstalled bool
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
	//add error check
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
	depPkgs := strings.Fields(p)
	for _, pkgs := range depPkgs {
		err = r.Command("sudo", "yum", "install", "-y", pkgs)
	}
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
	if !r.rvmInstalled {
		err := r.EnsureRvmInstalled()
		if err != nil {
			return err
		}
		r.rvmInstalled = true
	}
	err := r.Command("gem", "install", p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) EnsureRvmInstalled() error {

	r.InstallPackage(rvmDependencies)
	getRbenv := "curl -sL https://github.com/rbenv/rbenv-installer/raw/master/bin/rbenv-installer | bash -"
	r.Command("bash", "-c", getRbenv)
	setBashrc := `echo 'export PATH="$HOME/.rbenv/bin:$PATH"' >> $HOME/.bashrc && echo 'eval "$(rbenv init -)"' >> $HOME/.bashrc && source $HOME/.bashrc`
	var match bool = false
	if !match {
		r.Command("bash", "-c", setBashrc)
	}
	installRbenv := `rbenv install 2.7.0 && rbenv global 2.7.0`
	r.Command("bash", "-c", installRbenv)

	if r.Error != nil {
		fmt.Errorf("error installing rvm")
	}

	return nil
}

//Package server automates the installation of the Jekyll static website generator
package server

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

//Install dependencies for jekyll, apache and lets encrypt
const jekyllDep string = "gcc-c++ patch readline readline-devel zlib zlib-devel libffi-devel openssl-devel make bzip2 autoconf automake libtool bison sqlite-devel curl git-core"
const apacheDep string = "httpd"
const certbotDep string = "certbot python2-certbot-apache"

// Runner is a runner object with exportable and unexportable names
type Runner struct {
	History        []string
	Output         string
	Error          error
	GetRbenv       string
	SetBashrc      string
	InstallRbenv   string
	dryRun         bool
	yumUpdated     bool
	rbenvInstalled bool
}

// NewRunner returns a referecne to Runner
func NewRunner() *Runner {
	return &Runner{}
}

// NewTestRunner returns an instance of Runner for testing
func NewTestRunner() *Runner {
	r := NewRunner()
	r.dryRun = true //toggle to false to execute code locally
	return r
}

// Command reads from runner and takes one or more commandline arguments
func (r *Runner) Command(command string, args ...string) error {
	r.History = append(r.History, fmt.Sprintf("%s %s", command, strings.Join(args, " ")))
	if r.dryRun {
		return nil
	}

	output, err := exec.Command(command, args...).CombinedOutput()
	log.Println(string(output))
	if err != nil {
		return fmt.Errorf("failed to run '%s %s': %w", command, strings.Join(args, " "), err)
	}
	return nil
}

// InstallPackage reads from runner and installs dependency packages
func (r *Runner) InstallPackage(p string) error {
	err := r.EnsureYumUpdated() // update yum before installing packages
	if err != nil {
		return err
	}
	depPkgs := strings.Fields(p) // Loop through dependency packages
	for _, pkgs := range depPkgs {
		err = r.Command("sudo", "yum", "install", "-y", pkgs)
	}
	return err
}

// EnsureYumUpdated updates yum packages
func (r *Runner) EnsureYumUpdated() error {
	if r.yumUpdated {
		return nil
	}
	r.yumUpdated = true
	return r.Command("sudo", "yum", "update", "-y")
}

// InstallGem reads from runner, checks if dependencies
// are installed, and completes the install of jekyll and bundler
func (r *Runner) InstallGem(p string) error {
	r.EnsureRbenvInstalled()
	gemPath := "$HOME/.rbenv/shims/gem install " + p
	return r.Command("bash", "-c", gemPath)
}

// EnsureRbenvInstalled reads from runner and
// installs dependencies
func (r *Runner) EnsureRbenvInstalled() error {
	if r.rbenvInstalled {
		return nil
	}
	r.InstallPackage(jekyllDep)
	r.InstallPackage(apacheDep)
	r.InstallPackage(certbotDep)
	r.rbenvInstalled = true

	r.GetRbenv = "curl -sL https://github.com/rbenv/rbenv-installer/raw/master/bin/rbenv-installer | bash -"
	r.Command("bash", "-c", r.GetRbenv)

	r.SetBashrc = `echo 'export PATH="$HOME/.rbenv/bin:$PATH"' >> $HOME/.bashrc && echo 'eval "$(rbenv init -)"' >> $HOME/.bashrc && source $HOME/.bashrc`
	r.Command("bash", "-c", r.SetBashrc)

	r.InstallRbenv = "$HOME/.rbenv/bin/rbenv install 2.7.0 && $HOME/.rbenv/bin/rbenv global 2.7.0"
	r.Command("bash", "-c", r.InstallRbenv)

	return nil
}

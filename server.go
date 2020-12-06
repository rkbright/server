package server

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const jekyllDep string = "gcc-c++ patch readline readline-devel zlib zlib-devel libffi-devel openssl-devel make bzip2 autoconf automake libtool bison sqlite-devel curl git-core"
const apacheDep string = "httpd"
const certbotDep string = "certbot python2-certbot-apache"
<<<<<<< HEAD
=======

//group packages
//add doc comments
//look at comments on pkg.go.dev
>>>>>>> cccd7aae027d54b4e41f4d1a096092a16a9e5826

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
<<<<<<< HEAD
=======
	//add integration test, use a build tag
>>>>>>> cccd7aae027d54b4e41f4d1a096092a16a9e5826
	output, err := exec.Command(command, args...).CombinedOutput()
	log.Println(string(output))
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
	return err
}

func (r *Runner) EnsureYumUpdated() error {
	if r.yumUpdated {
		return nil
	}
	r.yumUpdated = true
	return r.Command("sudo", "yum", "update", "-y")
}

func (r *Runner) InstallGem(p string) error {
	r.EnsureRbenvInstalled()
	gemPath := "$HOME/.rbenv/shims/gem install " + p
	return r.Command("bash", "-c", gemPath)
}

func (r *Runner) EnsureRbenvInstalled() error {
	if r.rbenvInstalled {
		return nil
	}
	r.InstallPackage(jekyllDep)
	r.InstallPackage(apacheDep)
	r.InstallPackage(certbotDep)
	r.rbenvInstalled = true

<<<<<<< HEAD
	r.InstallPackage(jekyllDep)
	r.InstallPackage(apacheDep)
	r.InstallPackage(certbotDep)
	getRbenv := "curl -sL https://github.com/rbenv/rbenv-installer/raw/master/bin/rbenv-installer | bash -"
	r.Command("bash", "-c", getRbenv)
	// setBashrc := `echo 'export PATH="$HOME/.rbenv/bin:$PATH"' >> $HOME/.bashrc && echo 'eval "$(rbenv init -)"' >> $HOME/.bashrc && source $HOME/.bashrc`
	// r.Command("bash", "-c", setBashrc)
	installRbenv := `$HOME/.rbenv/bin/rbenv install 2.7.0 && $HOME/.rbenv/bin/rbenv global 2.7.0`
	r.Command("bash", "-c", installRbenv)
=======
	r.GetRbenv = "curl -sL https://github.com/rbenv/rbenv-installer/raw/master/bin/rbenv-installer | bash -"
	err := r.Command("bash", "-c", r.GetRbenv)
	if err != nil {
		return err
	}
>>>>>>> cccd7aae027d54b4e41f4d1a096092a16a9e5826

	r.SetBashrc = `echo 'export PATH="$HOME/.rbenv/bin:$PATH"' >> $HOME/.bashrc && echo 'eval "$(rbenv init -)"' >> $HOME/.bashrc && source $HOME/.bashrc`
	err = r.Command("bash", "-c", r.SetBashrc)
	if err != nil {
		return err
	}

	r.InstallRbenv = `$HOME/.rbenv/bin/rbenv install 2.7.0 && $HOME/.rbenv/bin/rbenv global 2.7.0`
	err = r.Command("bash", "-c", r.InstallRbenv)
	if err != nil {
		return err
	}
	return nil
}

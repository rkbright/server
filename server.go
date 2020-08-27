package server

import (
	"fmt"
	"os/exec"
	"strings"
)

const rvmDependencies string = "httpd certbot python2-certbot-apache gcc-c++ patch readline readline-devel zlib zlib-devel libffi-devel openssl-devel make bzip2 autoconf automake libtool bison sqlite-devel"

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
	getKey1 := "curl -sSL https://rvm.io/mpapis.asc | gpg2 --import -"
	r.Command("bash", "-c", getKey1)

	getKey2 := "curl -sSL https://rvm.io/pkuczynski.asc | gpg2 --import -"
	r.Command("bash", "-c", getKey2)

	getRvm := "curl -L get.rvm.io | bash -s stable"
	r.Command("bash", "-c", getRvm)

	r.Command("source", "$HOME/.rvm/scripts/rvm && exec bash")

	r.Command("rvm", "reload")
	r.Command("rvm", "requirements", "run")
	r.Command("rvm", "list", "known")
	r.Command("rvm", "install", "2.7")
	r.Command("rvm", "list")
	r.Command("rvm", "alias", "create", "default", "2.7")
	setRvmBashProfile := `echo -e "\n#set rvm\nif test -f ~/.rvm/scripts/rvm; then\n[ "$(type -t rvm)" = "function" ] || source ~/.rvm/scripts/rvm\nfi" >> ~/.bash_profile`
	setRvmBashrc := `echo -e "\n#set rvm\nif test -f ~/.rvm/scripts/rvm; then\n[ "$(type -t rvm)" = "function" ] || source ~/.rvm/scripts/rvm\nfi" >> ~/.bashrc`
	r.Command("bash", "-c", setRvmBashProfile)
	r.Command("bash", "-c", setRvmBashrc)
	r.Command("exec", "bash")
	r.Command("rvm", "default")

	if r.Error != nil {
		fmt.Errorf("error installing rvm")
	}

	return nil
}

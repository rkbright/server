package server

import (
	"fmt"
	"os/exec"
	"strings"
)

type Runner struct {
	History      []string
	Output       string
	dryRun       bool
	yumUpdated   bool
	rvmInstalled bool
	installed    bool
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

func (r *Runner) curlRvm(p string) error {
	err := r.Command("curl", p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) sourceRvm(p string) error {
	err := r.Command("source", p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) installRvm(p string) error {
	err := r.Command("rvm", p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) EnsureRvmInstalled() error {

	var rvmPackages string = "epel-release gcc-c++ patch readline readline-devel zlib zlib-devel libffi-devel openssl-devel make bzip2 autoconf automake libtool bison sqlite-devel"
	var gpgKey string = "-sSL https://rvm.io/mpapis.asc | gpg2 --import -"
	var gpgKeySync string = "-sSL https://rvm.io/pkuczynski.asc | gpg2 --import -"
	var getRVM string = "-L get.rvm.io | bash -s stable"
	var sourceRVM string = "/etc/profile.d/rvm.sh"
	var rvmLoad string = "reload"
	var rvmReq string = "requirements run"
	var installRvm string = "install 2.7"
	var useRvm string = "use 2.7 --default"

	err := r.InstallPackage(rvmPackages)
	if err != nil {
		return err
	}

	err = r.curlRvm(gpgKey)
	if err != nil {
		return err
	}

	err = r.curlRvm(gpgKeySync)
	if err != nil {
		return err
	}

	err = r.curlRvm(getRVM)
	if err != nil {
		return err
	}

	err = r.sourceRvm(sourceRVM)
	if err != nil {
		return err
	}

	err = r.installRvm(rvmLoad)
	if err != nil {
		return err
	}

	err = r.installRvm(rvmReq)
	if err != nil {
		return err
	}

	err = r.installRvm(installRvm)
	if err != nil {
		return err
	}

	err = r.installRvm(useRvm)
	if err != nil {
		return err
	}

	return nil
}

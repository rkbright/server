package server_test

import (
	"server"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInstallPackage(t *testing.T) {
	t.Parallel()
	r := server.NewTestRunner()
	err := r.InstallPackage("git")
	if err != nil {
		t.Fatal(err)
	}
	err = r.InstallPackage("wget")
	if err != nil {
		t.Fatal(err)
	}
	wantHistory := []string{
		"sudo yum update -y",
		"sudo yum install -y git",
		"sudo yum install -y wget",
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

func TestInstallGem(t *testing.T) {
	t.Parallel()
	r := server.NewTestRunner()
	err := r.InstallGem("jekyll")
	if err != nil {
		t.Fatal(err)
	}
	err = r.InstallGem("bundler")
	if err != nil {
		t.Fatal(err)
	}
	wantHistory := []string{
		"sudo yum update -y",
		"sudo yum install -y gcc-c++",
		"sudo yum install -y patch",
		"sudo yum install -y readline",
		"sudo yum install -y readline-devel",
		"sudo yum install -y zlib",
		"sudo yum install -y zlib-devel",
		"sudo yum install -y libffi-devel",
		"sudo yum install -y openssl-devel",
		"sudo yum install -y make",
		"sudo yum install -y bzip2",
		"sudo yum install -y autoconf",
		"sudo yum install -y automake",
		"sudo yum install -y libtool",
		"sudo yum install -y bison",
		"sudo yum install -y sqlite-devel",
		"sudo yum install -y curl",
		"sudo yum install -y git-core",
		"sudo yum install -y httpd",
		"sudo yum install -y certbot",
		"sudo yum install -y python2-certbot-apache",
		"bash -c " + r.GetRbenv,
		`bash -c ` + r.SetBashrc,
		"bash -c " + r.InstallRbenv,
		"bash -c $HOME/.rbenv/shims/gem install jekyll",
		"bash -c $HOME/.rbenv/shims/gem install bundler",
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

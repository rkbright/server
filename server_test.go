package server_test

import (
	"server"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInstallPackage(t *testing.T) {
	t.Parallel()
	r := server.NewTestRunner()
	err := r.InstallPackage("python")
	if err != nil {
		t.Fatal(err)
	}
	err = r.InstallPackage("java")
	if err != nil {
		t.Fatal(err)
	}
	wantHistory := []string{
		"sudo yum update -y",
		"sudo yum install -y python",
		"sudo yum install -y java",
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

func TestInstallGem(t *testing.T) {
	t.Parallel()
	r := server.NewTestRunner()
	err := r.InstallGem("bundler")
	if err != nil {
		t.Fatal(err)
	}
	err = r.InstallGem("jekyll")
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
		"curl -sSL https://rvm.io/mpapis.asc | gpg2 --import - ",
		"curl -sSL https://rvm.io/pkuczynski.asc | gpg2 --import - ",
		"curl -L get.rvm.io | bash -s stable ",
		"source /etc/profile.d/rvm.sh ",
		"rvm reload ",
		"rvm requirements run ",
		"rvm install 2.7 ",
		"rvm use 2.7 --default ",
		"gem install bundler",
		"gem install jekyll",
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

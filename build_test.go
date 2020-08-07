// Package to test build.go file
package thing_test

import (
	"testing"
	"thing"

	"github.com/google/go-cmp/cmp"
)

func TestInstallPackage(t *testing.T) {
	t.Parallel()
	r := thing.NewTestRunner()

	err := r.InstallPackage("python")
	if err != nil {
		t.Fatal(err)
	}
	err = r.InstallPackage("java")
	if err != nil {
		t.Fatal(err)
	}
	wantHistory := []string{
		"yum update -y",
		"yum install -y python",
		"yum install -y java",
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

func TestInstallGem(t *testing.T) {
	t.Parallel()
	r := thing.NewTestRunner()
	err := r.InstallGem("bundler")
	if err != nil {
		t.Fatal(err)
	}
	err = r.InstallGem("jekyll")
	if err != nil {
		t.Fatal(err)
	}
	wantHistory := []string{
		"yum update -y",
		"yum install -y epel-release",
		"yum install -y ruby",
		"gem install bundler",
		"gem install jekyll",
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

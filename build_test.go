// Package to test build.go file
package thing_test

import (
	"testing"
	"thing"

	"github.com/google/go-cmp/cmp"
)

// for yum packages, test the user has an id of 0 (root)
func TestUserCommand(t *testing.T) {
	t.Parallel()
	r := thing.TestNewRunner()
	err := r.Command("id", "-u")

	wantHistory := []string{"id -u"}
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

func TestInstallPackage(t *testing.T) {
	t.Parallel()
	r := thing.TestNewRunner()

	err := r.InstallPackage([]string{"update", "epel-release", "ruby"})
	if err != nil {
		t.Fatal(err)
	}
	wantHistory := []string{
		"yum update -y",
		"yum install -y epel-release",
		"yum install -y ruby",
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

func TestInstallGem(t *testing.T) {
	t.Parallel()
	r := thing.TestNewRunner()

	err := r.InstallGems([]string{"bundler", "jekyll"})
	if err != nil {
		t.Fatal(err)
	}
	wantHistory := []string{
		"gem install bundler",
		"gem install jekyll",
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

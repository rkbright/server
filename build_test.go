// Package to test build.go file
package thing_test

import (
	"testing"
	"thing"

	"github.com/google/go-cmp/cmp"
)

func TestCommand(t *testing.T) {
	t.Parallel()

	r := thing.NewTestRunner()
	err := r.Command("echo", "You successfully ran a Linux command from Go!!!")

	wantHistory := []string{"echo You successfully ran a Linux command from Go!!!"}
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

func TestYumUpdate(t *testing.T) {
	t.Parallel()
	r := thing.NewTestRunner()

	err := r.UpdateYum()
	if err != nil {
		t.Fatal(err)
	}

	wantHistory := []string{"yum update -y"}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

func TestInstallPackages(t *testing.T) {
	t.Parallel()
	r := thing.NewTestRunner()

	err := r.InstallPackages([]string{"epel-release", "ruby"})
	if err != nil {
		t.Fatal(err)
	}
	wantHistory := []string{
		"yum install -y epel-release",
		"yum install -y ruby",
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

func TestInstallGems(t *testing.T) {
	t.Parallel()
	r := thing.NewTestRunner()

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

func TestCheckInstalledPackages(t *testing.T) {
	t.Parallel()
	r := thing.NewTestRunner()

	_ = r.IsInstalled("ruby")
	wantHistory := []string{
		"rpm -q ruby",
	}
	if !cmp.Equal(wantHistory, r.History) {
		t.Fatal(cmp.Diff(wantHistory, r.History))
	}
}

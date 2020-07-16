// Package to test build.go file
package thing_test

import (
	"testing"
	"thing"
)

func TestCommand(t *testing.T) {
	t.Parallel()

	r := thing.NewRunner()
	r.Test = true
	err := r.Command("echo", "You successfully ran a Linux command from Go!!!")

	want := "echo You successfully ran a Linux command from Go!!!"
	if err != nil {
		t.Fatal(err)
	}
	if want != r.CmdLine {
		t.Errorf("got %q, want %q", want, r.CmdLine)
	}
}

func TestYumUpdate(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	err := r.YumUpdate()
	if err != nil {
		t.Fatal(err)
	}

	wantCmd := "yum update -y"
	if wantCmd != r.CmdLine {
		t.Errorf("got %q, want %q", wantCmd, r.CmdLine)
	}
}

func TestYumInstall(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	err := r.YumInstall()
	if err != nil {
		t.Fatal(err)
	}
	wantCmd := "yum install epel-release -y"
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

func TestRubyInstall(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	err := r.RubyInstall()
	if err != nil {
		t.Fatal(err)
	}

	wantCmd := "yum install ruby -y"
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

func TestJekyllInstall(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	err := r.JekyllInstall()
	if err != nil {
		t.Fatal(err)
	}

	wantCmd := "gem install jekyll"
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

func TestBundlerInstall(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	err := r.BundlerInstall()
	if err != nil {
		t.Fatal(err)
	}
	wantCmd := "gem install bundler"
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

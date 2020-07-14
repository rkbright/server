// Package to test build.go file
package thing_test

import (
	"testing"
	"thing"
)

func TestCommand(t *testing.T) {
	t.Parallel()
	want := "You successfully ran a Linux command from Go!!!\n"
	got, err := thing.NewRunner().Command("echo", "You successfully ran a Linux command from Go!!!")
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestYumUpdate(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = false
	r.YumUpdate()

	wantCmd := "yum update -y"
	got := r.CmdLine

	if wantCmd != got {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

func TestYumInstall(t *testing.T) {
	t.Parallel()
	wantCmd := "yum install epel-release -y"
	r := thing.NewRunner()
	r.Test = true
	r.YumInstall()
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

func TestRubyInstall(t *testing.T) {
	t.Parallel()
	wantCmd := "yum install ruby -y"
	r := thing.NewRunner()
	r.Test = true
	r.RubyInstall()
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

func TestJekyllInstall(t *testing.T) {
	t.Parallel()
	wantCmd := "gem install jekyll"
	r := thing.NewRunner()
	r.Test = true
	r.JekyllInstall()
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

func TestBundlerInstall(t *testing.T) {
	t.Parallel()
	wantCmd := "gem install bundler"
	r := thing.NewRunner()
	r.Test = true
	r.BundlerInstall()
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

func TestRubyVersion(t *testing.T) {
	t.Parallel()
	//ruby versions will change overtime, therefore the test only verifies
	//that a ruby version exists on the server and not for a specific version
	wantCmd := "ruby 2.6"
	r := thing.NewRunner()
	r.Test = true
	r.RubyVersion()
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

func TestJekyllVersion(t *testing.T) {
	t.Parallel()
	wantCmd := "jekyll 4.*"
	r := thing.NewRunner()
	r.Test = true
	r.JekyllVersion()
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

func TestBundlerVersion(t *testing.T) {
	t.Parallel()
	wantCmd := "Bundler 2.*"
	r := thing.NewRunner()
	r.Test = true
	r.BundlerVersion()
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

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
	r.Test = true

	wantCmd := "yum update -y"
	got := r.YumUpdate()

	if wantCmd != got {
		t.Fatalf("want command %q, got %q", wantCmd, got)
	}
}

func TestYumInstall(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	wantCmd := "yum install epel-release -y"
	got := r.YumInstall()

	if wantCmd != got {
		t.Fatalf("want command %q, got %q", wantCmd, got)
	}
}

func TestRubyInstall(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	wantCmd := "yum install ruby -y"
	got := r.RubyInstall()
	if wantCmd != got {
		t.Fatalf("want command %q, got %q", wantCmd, got)
	}
}

func TestJekyllInstall(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	wantCmd := "gem install jekyll"
	got := r.JekyllInstall()
	if wantCmd != got {
		t.Fatalf("want command %q, got %q", wantCmd, got)
	}
}

func TestBundlerInstall(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	wantCmd := "gem install bundler"
	got := r.BundlerInstall()
	if wantCmd != got {
		t.Fatalf("want command %q, got %q", wantCmd, got)
	}
}

func TestRubyVersion(t *testing.T) {
	t.Parallel()
	//ruby versions will change overtime, therefore the test only verifies
	//that a ruby version exists on the server and not for a specific version
	r := thing.NewRunner()
	r.Test = true

	wantCmd := "ruby 2.6"
	got := r.RubyVersion()
	if wantCmd != got {
		t.Fatalf("want command %q, got %q", wantCmd, got)
	}
}

func TestJekyllVersion(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	wantCmd := "jekyll 4.*"
	got := r.JekyllVersion()
	if wantCmd != got {
		t.Fatalf("want command %q, got %q", wantCmd, got)
	}
}

func TestBundlerVersion(t *testing.T) {
	t.Parallel()
	r := thing.NewRunner()
	r.Test = true

	wantCmd := "Bundler 2.*"
	got := r.BundlerVersion()
	if wantCmd != got {
		t.Fatalf("want command %q, got %q", wantCmd, got)
	}
}

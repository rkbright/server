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
	wantCmd := "yum update -y"
	r := thing.NewRunner()
	r.Test = true
	r.YumUpdate()
	if wantCmd != r.CmdLine {
		t.Fatalf("want command %q, got %q", wantCmd, r.CmdLine)
	}
}

// Package to test build.go file
package build_test

import (
	"build"
	"testing"
)

func TestBuild(t *testing.T) {
	got := build.Runfile()
	want := "You successfully ran a Linux command from Go!!!" + "\n"
	test := "you successfully ran a linux command from go!!!" + "\n"

	if got != want {
		t.Errorf("got %q, want %q, test %q", got, want, test)
	}
}

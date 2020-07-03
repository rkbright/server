// Package to test build.go file
package build_test

import (
	"build"
	"testing"
)

// TestBuild function
func TestBuild(t *testing.T) {
	got := build.Runfile()
	want := "You successfully ran a Linux command from Go!!!" + "\n"
	test2 := "you successfully ran a linux command from go!!!" + "\n"
	if got != want {
		t.Errorf("got %q, want %q, test %q", got, want, test2)
	}
}

// TestNative function
func TestNative(t *testing.T) {
	got := build.Native()
	want := 0

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

// Package to test build.go file
package build_test

import (
	"src/build"
	"testing"
)

// TestBuild function
func TestBuild(t *testing.T) {
	got, err := build.Runfile()
	want := "You successfully ran a Linux command from Go!!!\n"
	test2 := "you successfully ran a linux command from go!!!\n"
	if want != got {
		t.Errorf("got %q, want %q, test %q, error %q", got, want, test2, err)
	}
}

// TestNative function
func TestNative(t *testing.T) {
	got, err := build.Native()
	want := "You successfully ran a Linux command from Go!!!\n"
	if want != got {
		t.Errorf("want %q, got %q, error %q", want, got, err)
	}
}

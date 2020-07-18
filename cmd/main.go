package main

import (
	"thing"
)

func main() {
	r := thing.NewRunner()
	r.InstallPackages([]string{"epel-release"})
	r.UpdateYum()
	r.InstallPackages([]string{"ruby"})
	r.InstallGems([]string{"jekyll", "bundler"})
	r.CheckInstalledPackages([]string{"ruby", "bundler", "gem"})
}

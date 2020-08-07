package main

import (
	"thing"
)

func main() {
	r := thing.NewRunner()
	r.InstallPackage("epel-release")
	r.InstallPackage("ruby")
	r.InstallGem("bundler")
	r.InstallGem("jekyll")
}

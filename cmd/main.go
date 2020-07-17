package main

import (
	"thing"
)

func main() {
	r := thing.NewRunner()
	r.InstallPackages(["epel"])
	r.UpdateYum()
	r.InstallPackages(["ruby"])
	r.InstallGems(["bundler", "jekyll"])
}

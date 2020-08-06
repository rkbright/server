package main

import (
	"thing"
)

func main() {

	r := thing.NewRunner()
	r.InstallPackage([]string{"epel-release", "ruby"})
	r.InstallGems([]string{"bundler", "jekyll"})

}

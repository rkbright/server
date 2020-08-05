package main

import "thing"

func main() {

	r := thing.NewRunner()
	r.InstallPackage([]string{"update", "epel-release", "ruby"})
	r.InstallGems([]string{"bundler", "jekyll"})

}

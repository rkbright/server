package main

import (
	"thing"
)

func main() {
	r := thing.NewRunner()
	r.InstallGem("bundler")
	r.InstallGem("jekyll")
}

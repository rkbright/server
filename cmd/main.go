package main

import (
	"server"
)

// Installs jekyll and bundler
func main() {
	r := server.NewRunner()
	r.InstallGem("jekyll")
	r.InstallGem("bundler")

}

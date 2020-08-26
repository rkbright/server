package main

import (
	"server"
)

func main() {
	r := server.NewRunner()
	r.InstallGem("bundler")
	r.InstallGem("jekyll")

}

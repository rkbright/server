package main

import (
	"server"
)

func main() {
	r := server.NewRunner()
	r.InstallGem("jekyll")
	r.InstallGem("bundler")

}

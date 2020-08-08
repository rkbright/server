package main

import (
	"log"
	"server"
)

func main() {
	r := server.NewRunner()
	// r.InstallGem("bundler")
	// r.InstallGem("jekyll")
	err := r.Command("date")
	if err != nil {
		log.Fatal(err)
	}
}

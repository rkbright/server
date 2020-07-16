package main

import (
	"thing"
)

func main() {
	thing.NewRunner().YumUpdate()
	thing.NewRunner().YumInstall()
	thing.NewRunner().RubyInstall()
	thing.NewRunner().BundlerInstall()
	thing.NewRunner().JekyllInstall()
}

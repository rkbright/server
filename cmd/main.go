package main

import "thing"

func main() {
	YumUpdate()
	YumInstall("epel-release")
	YumInstall("ruby")
	GemInstall("jekyll")
	GemInstall("bundler")
}

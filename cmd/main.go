package main

import "thing"

func main() {
	r := thing.NewRunner()
	r.InstallPackages([]string{"epel-release"})
	r.UpdateYum()
	r.InstallPackages([]string{"ruby"})
	r.InstallGems([]string{"jekyll", "bundler"})
	r.CheckInstalledPackages([]string{"ruby", "bundler", "gem"})
	r.CheckPackageExists([]string{"ruby", "bundler", "gem"})
}

/*
Notes:
When testing on GCP linux box I noted the following:
1. need to cross compile to linux/amd64
2. do not assume wget or curl are available on the machine
3. do not assume git is available on the machine
4. commands did not run on linux machine

*/

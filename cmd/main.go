package main

import "thing"

func main() {

	packInstall := make([]string, 5)
	packInstall[0] = "ruby"
	packInstall[1] = "jekyll"
	packInstall[2] = "bundler"
	packInstall[3] = "gem"
	packInstall[4] = "epel-release"

	r := thing.NewRunner()
	r.InstallPackages(packInstall[4:])
	r.UpdateYum()
	r.InstallPackages(packInstall[0:])
	r.InstallGems(packInstall[1:2])
	r.CheckInstalledPackages(packInstall[:])
	r.CheckPackageExists(packInstall[:])
}

/*
Notes:
When testing on GCP linux box I noted the following:
1. need to cross compile to linux/amd64
2. do not assume wget or curl are available on the machine
3. do not assume git is available on the machine
4. commands did not run on linux machine

*/

[![Go Report Card](https://goreportcard.com/badge/github.com/rkbright/jekyllizer)](https://goreportcard.com/report/github.com/rkbright/jekyllizer)

# How can I use `Jekyllizer`?

**PLEASE NOTE**: this package was build for yum-based linux distributions and was tested on RHEL and CentOS 7 operating systems. I do plan on extending the package to Debian-based systems in the future. Please feel free to add the functionality! 

**IN ADDITION**: you will need to run the program from an account with elevated privileges, i.e., able to run sudo commands. Do not run from the root account.  

1. **Install git** 

    `sudo yum install git -y`

2. **Install Go**

    Follow the installation steps on the [Go download and install site](https://golang.org/doc/install).

        1. Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go. For example, run the following as root or through sudo:

        tar -C /usr/local -xzf *.tar.gz

        2. Add /usr/local/go/bin to the PATH environment variable.

        You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

        export PATH=$PATH:/usr/local/go/bin

        Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

        3. Verify that you've installed Go by opening a command prompt and typing the following command:

        $ go version

        Confirm that the command prints the installed version of Go.

3. **Then clone the repo into your $HOME directory**

    `git clone https://github.com/rkbright/jekyllizer.git`


4. Change directories 

    `cd jekyllizer`

    Then run the program `go run cmd/main.go`

That's it! It takes a few minutes to run yum update and download the dependencies, especially for ephemeral VMs that are newly provisioned. So feel free to grab a cup of your favorite beverage.

# What is `Jekyllizer`? 

`Jekyllizer` is a Go package that automates the build process to support Jekyll-based static websites.  

What will be installed?
* ruby version 2.7 (you can update the version in the `server.go` file)
* jekyll 
* bundler 
* gem
* apache 
* lets encrypt 

# What is Jekyll?

[Jekyll](https://jekyllrb.com/) is a simple, blog-aware, static site generator for personal, project, or organization sites. 

# GoDoc

Link to [GoDoc page](https://godoc.org/github.com/rkbright/jekyllizer)

# Other Resources

I know of only one other repo that tries to automate the Jekyll install 

Parker Moore: https://github.com/parkr/jekyll-build-server 
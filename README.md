# How can I use `Jekyll server build`?

**PLEASE NOTE**: this package was build for yum-based linux distributions and was tested on RHEL and CentOS 7 operating systems. I do plan on extending the package to Debian-based systems in the future. Please feel free to add the functionality! 

1. **Install git** 

    `sudo yum install git -y`

2. **Install Go**

    Follow the installation steps on the [Go download and install site](https://golang.org/doc/install).

        1. Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go. For example, run the following as root or through sudo:

        `tar -C /usr/local -xzf *go1.*15.6.linux-amd64*.tar.gz*`

        2. Add /usr/local/go/bin to the PATH environment variable.

        You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

        `export PATH=$PATH:/usr/local/go/bin` 

        >Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.

        3. Verify that you've installed Go by opening a command prompt and typing the following command:

        `$ go version`

        Confirm that the command prints the installed version of Go.

3. **Then clone the repo into your $HOME directory**

    `git clone https://github.com/rkbright/jekyll-server-build.git`

# What is `Jekyll server build`? 

`Jekyll server build` is a Go package that automates the build process to support Jekyll-based static websites.  

What will be installed?
* ruby version 2.7 (you can update the version in the `server.go` file)
* bundler 
* jekyll 
* gem
* apache 
* lets encrypt 

# What is Jekyll?

[Jekyll](https://jekyllrb.com/) is a simple, blog-aware, static site generator for personal, project, or organization sites. 

# Other Resources

I know of only one other repo that tries to automate the Jekyll install 

Parker Moore: https://github.com/parkr/jekyll-build-server 
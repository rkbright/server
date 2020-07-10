!#/bin/bash/

#run yum update and install epel package 
sudo yum update -y 
sudo yum install epel-release -y

#install ruby
sudo yum install ruby -y

#install gem
gem install jekyll 

#install bundler 
gem install bundler

sudo apt-get install git
sudo apt-get install build-essential python-dev # for mongoengine
sudo apt-get install python-setuptools
sudo apt-get install openssh-server
sudo apt-get install nginx
sudo apt-get install supervisor
sudo apt-get install npm # just for coffee script
sudo apt-get install gem # just for SASS

sudo apt-key adv --keyserver keyserver.ubuntu.com --recv 7F0CEB10
sudo sh -c "echo 'deb http://downloads-distro.mongodb.org/repo/ubuntu-upstart dist 10gen' > /etc/apt/sources.list.d/10gen.list"
sudo add-apt-repository ppa:webupd8team/sublime-text-2

sudo apt-get update
sudo apt-get install sublime-text
sudo apt-get install mongodb-10gen

sudo easy_install virtualenv
sudo npm install -g coffee-script
sudo gem install sass


#sudo apt-get install mysql-server  # Just for development machines. AWS will give us our database. 


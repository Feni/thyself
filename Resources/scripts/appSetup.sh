sudo mkdir /var/www
sudo mkdir /var/cache/nginx
sudo git clone https://Feni@bitbucket.org/Feni/thyself.git /var/www/thyself
cd /var/www/thyself
sudo chown -R feni /var/www/thyself  # Where feni is my username
git config --global core.excludesfile '/var/www/thyself/.gitignore'
cp /var/www/thyself/server/.vimrc ~/.vimrc
sudo virtualenv venv --no-site-packages
sudo cp /var/www/thyself/server/thyself.conf /etc/nginx/sites-available/thyself.conf
sudo ln -s /etc/nginx/sites-available/thyself.conf /etc/nginx/sites-enabled/
sudo cp /var/www/thyself/server/nginx.conf /etc/nginx/nginx.conf
sudo rm /etc/nginx/sites-enabled/default
sudo nginx -t
sudo service nginx reload

#!/bin/bash 






sudo apt-get install

sudo apt-get install ngnix -y

sudo systemctl start ngnix

sudo systemctl enable ngnix

cp hello.html /var/www/html

sudo systemctl restart ngnix 

echo "Devborad in running on port 80"

#!/bin/bash

sudo docker run -i -t -p 8080:8080 --rm tcotav/glockv /bin/bash
#MYSQL_ROOT=
#
## pull the container
#docker pull mysql
#
## ref: https://hub.docker.com/_/mysql/
## start mysql
##docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT -d mysql:5.7
#docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=bYt0cun3pev8cH -d mysql:5.7
#
## run app container
##docker run --name some-app --link some-mysql:mysql -d application-that-uses-mysql
#
#
## this gets you a mysql client prompt
##docker run -it --link some-mysql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ENV_MYSQL_ROOT_PASSWORD"'
#sudo docker run -it --link some-mysql:mysql --rm mysql sh -c 'exec mysql -h"$MYSQL_PORT_3306_TCP_ADDR" -P"$MYSQL_PORT_3306_TCP_PORT" -uroot -p"$MYSQL_ROOT"'
#
## run the app
#sudo docker run -i -t -p 8080:8080 --link some-mysql:mysql --rm tcotav/glockv /bin/bash

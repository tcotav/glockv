#!/bin/bash

# on ubuntu

sudo apt-get update
sudo apt-get install sqlite3 libsqlite3-dev -y

# assumes you already ran setup
sqlite3 geoloc.db < sql/create.sql


# we do this so we have a self-contained package
# export GOPATH=`pwd` 
go get github.com/Sirupsen/logrus
go get github.com/mattn/go-sqlite3
go get github.com/TomiHiltunen/geohash-golang

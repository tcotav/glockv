#!/bin/bash

# on ubuntu

sudo apt-get update
sudo apt-get install sqlite3 libsqlite3-dev -y

# assumes you already ran setup
sqlite3 geoloc.db < sql/create.sql


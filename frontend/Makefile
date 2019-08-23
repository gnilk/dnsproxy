#
# Frontend Makefile
#
# basic targets:
#   make bundle		- builds the client side bundle
#
# Special targets: 
#   make uidev      - starts the webpack dev server with hot-module replacement on port 8080, use this for front end development
#   make production - builds the bundle using webpack production configuraion
#
# Notes: 
#  webpack-server (uidev) does not create a bundle instead it keeps the bundle in memory for hot module replacement
#  the service (appservice) must be started from the 'app' folder to access the files, normally you can run '../../../bin/appservice'
#
#
BIN = $(GOPATH)/bin

NODE_BIN = $(shell npm bin)
NAME = app


VERSION=0


all:	bundle


bundle:	
	$(NODE_BIN)/webpack -p --progress --bail --config webpack.prod.js

production:	
	$(NODE_BIN)/webpack -p --progress --colors --bail --config webpack.prod.js

uidev:
	$(NODE_BIN)/webpack-dev-server --host 0.0.0.0 --inline --hot -d --progress --colors --config webpack.dev.js


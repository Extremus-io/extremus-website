#!/usr/bin/env bash

npm install -g bower

npm install -g polymer-cli

cd webapp

bower install

polymer build

#!/usr/bin/env bash

echo "Dependent libraries installation"
export GOPATH=./src

go get github.com/gookit/color

echo "Finish the program!"
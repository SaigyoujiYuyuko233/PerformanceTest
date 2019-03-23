#!/usr/bin/env bash

echo "Dependent libraries installation"
export GOPATH=./src

go get
go install

echo "Finish the program!"
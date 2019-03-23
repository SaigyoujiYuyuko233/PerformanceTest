#!/usr/bin/env bash

echo "Dependent libraries installation"

go get github.com/gookit/color
go install github.com/gookit/color

echo "Finish the program!"
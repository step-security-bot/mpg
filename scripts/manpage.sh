#!/bin/sh -e
rm -rf manpage
mkdir manpage
go run main.go man >manpage/mpg.1

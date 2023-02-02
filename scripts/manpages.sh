#!/bin/sh -e
rm -rf manpages
mkdir manpages
go run . man
mv *.1 manpages

#!/bin/bash

cmd="/Users/q187392/dev/gom/bin/g"

editCsv=./test.edit.csv
openCsv=./test.open.csv
jumpCsv=./test.jump.csv

tE1 () {
    printf "Running tE1\n"
    #$cmd -f $editCsv
    go run g.go -f $editCsv gostart
    return 0
}

tE1

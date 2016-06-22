#!/bin/bash

g () {
    file=$($GOBIN/g $1)
    if [ $? -eq 0 ]; then
        builtin cd $file
    fi
}



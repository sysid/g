#!/bin/bash

. $HOME/binx/myFunc.x

cmd="/Users/q187392/dev/gom/bin/g"

editCsv=./test.edit.csv
openCsv=./test.open.csv
jumpCsv=./test.jump.csv

tE1 () {
    printf "____Running tE1: key exists, but 2 entries in CSV\n"
    #$cmd -f $editCsv
    ref="/Users/q187392/binx/startGo.x"
    got=$(go run g.go -f $editCsv gostart)
    if [ $got = $ref ]; then
        Green "Pass: $got"
    else
        printf "Expected: %s\n" $ref
        printf "Got: %s\n" $got
        Red "NOK"
    fi
    return 0
}

tE2 () {
    printf "____Running tE2: key not in map.\n"
    ref=15
    got=$(go run g.go -f $editCsv gostartt 2>&1 | wc -l)
    if [ "$got" -eq "$ref" ]; then
        Green "Pass: $got"
    else
        printf "Expected: %s\n" $ref
        printf "Got: %s\n" $got
        Red "NOK"
    fi
    return 0
}

tE3 () {
    printf "____Running tE3: no key provided as parameter.\n"
    ref=15
    got=$(go run g.go -f $editCsv 2>&1 | wc -l)
    if [ $got = $ref ]; then
        Green "Pass: $got"
    else
        printf "Expected: %s\n" $ref
        printf "Got: %s\n" $got
        Red "NOK"
    fi
    return 0
}

tE4 () {
    printf "____Running tE4: check retCode, key not in map.\n"
    ref=1
    go run g.go -f $editCsv xxx >/dev/null 2>&1
    got=$?
    if [ $got -eq $ref ]; then
        Green "Pass: $got"
    else
        printf "Expected: %s\n" $ref
        printf "Got: %s\n" $got
        Red "NOK"
    fi
    return 0
}

tE5 () {
    printf "____Running tE5: check retcode, key exists\n"
    ref=0
    go run g.go -f $editCsv gostart > /dev/null 2>&1
    got=$?
    if [ $got = $ref ]; then
        Green "Pass: $got"
    else
        printf "Expected: %s\n" $ref
        printf "Got: %s\n" $got
        Red "NOK"
    fi
    return 0
}

tE6 () {
    printf "____Running tE6: check retcode, key exists, debugging enabled\n"
    ref=0
    go run g.go -f $editCsv -d gostart > /dev/null 2>&1
    got=$?
    if [ $got = $ref ]; then
        Green "Pass: $got"
    else
        printf "Expected: %s\n" $ref
        printf "Got: %s\n" $got
        Red "NOK"
    fi
    return 0
}

tE1
tE2
tE3
tE4
tE5
tE6

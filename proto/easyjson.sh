#!/bin/bash

go get github.com/mailru/easyjson && \
go install github.com/mailru/easyjson/...@latest

if ! which easyjson &> /dev/null
then
    exit 1
fi

rm -fr *_easyjson.go

for gofile in $(ls *.go | grep -v -e easyjson -e gen.go)
do
    echo ${gofile}
    easyjson -omit_empty ${gofile}
done

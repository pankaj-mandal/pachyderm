#!/bin/bash
set -xeuo pipefail

while true; do

    pachctl create repo data
    echo "version 1" | pachctl put file data@master:/myfile.txt
    curl -XPUT 'localhost:9002/repos/data/master/_mount?name=data&mode=ro'
    cat /pfs/data2/myfile.txt
    curl -XPUT 'localhost:9002/repos/data/master/_unmount?name=data'
    pachctl create branch data@v1 --head master
    curl -XPUT 'localhost:9002/repos/data/master/_mount?name=data&mode=ro'
    cat /pfs/data2/myfile.txt
    pachctl delete repo data

done

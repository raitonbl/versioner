#!/bin/bash
set -e

if [ "$1" == "get" ]; then
  eval "versioner $1 $2 $3 $4 $5"
elif [ "$1" == "set" ]; then
   eval "versioner $1 $2 $3 $4 $5"
elif [ "$1" == "set-stamped-version" ]; then
  eval "versioner $1 $2 $3 $4 $5"
else
    echo "command $1 not supported"
    exit 1
fi
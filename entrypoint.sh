#!/bin/bash

set -e

curl "https://s3.af-south-1.amazonaws.com/binaries.raitonbl.com/public/versioner/$1/versioner" --output versioner
chmod a+x versioner
mv versioner /usr/local/bin/versioner

ARGS="versioner $2 $3 $4 $5 $6"

eval "$ARGS"
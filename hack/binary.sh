#!/usr/bin/env bash

set -e

IMPORTPATH=github.com/frankgreco/fetakv
PROJECTPATH=$GOPATH/src/$IMPORTPATH
BINARIES=($(for i in $(ls -d $PROJECTPATH/cmd/*/); do echo ${i%%/} | awk -F "/" '{print $NF}'; done))

for i in "${BINARIES[@]}"
do
  rm -f ${GOPATH%%:*}/bin/$i

  go install $IMPORTPATH/cmd/$i

  if [ $? -eq 0 ]; then
    echo "Build successful. Binary located at ${GOPATH%%:*}/bin/$i"
  else
    echo "Build failed."
  fi
done

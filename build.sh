#!/bin/bash

echo "*************************************"
echo "* B U I L D   A P P L I C A T I O N *"
echo "*************************************"

GO_BUILD="CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ."

echo "run $GO_BUILD"
eval $GO_BUILD
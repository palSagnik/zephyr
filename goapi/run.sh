#!/bin/zsh

rm test

echo "Building the go-project"
go build -o test

echo "Running the project"
./test

echo "Deleting the binary"
rm test

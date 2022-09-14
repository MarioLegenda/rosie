#!/bin/bash
echo "Compiling - Mac - 64-bit"
# 64-bit
GOOS=darwin GOARCH=amd64 go build -o bin/rosie64mac

echo "Compiling - Linux - 64-bit"
# 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/rosie64linux
echo "Compiling - Linux - 32-bit"
# 32-bit
GOOS=linux GOARCH=386 go build -o bin/rosie32linux

echo "Compiling - Native"
go build -o bin/rosie
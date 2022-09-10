#!/bin/bash

echo "Compiling - Windows - 64-bit"
# 64-bit
GOOS=windows GOARCH=amd64 go build -o bin/simulator64.exe
echo "Compiling - Windows - 32-bit"
# 32-bit
GOOS=windows GOARCH=386 go build -o bin/simulator32.exe

echo "Compiling - Mac - 64-bit"
# 64-bit
GOOS=darwin GOARCH=amd64 go build -o bin/simulator64

echo "Compiling - Linux - 64-bit"
# 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/simulator64linux
echo "Compiling - Linux - 32-bit"
# 32-bit
GOOS=linux GOARCH=386 go build -o bin/simulator32linux
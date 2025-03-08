#!/bin/bash


# ensure go is installed

echo "Checking if Go is installed..."
if ! command -v go & > /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi


if ! command -v python3 & > /dev/null; then
    echo "Python3 is not installed. Please install Python3 and try again."
    exit 1
fi


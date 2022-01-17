#!/bin/bash

# Install Dependency
echo "Install Dependencies..."
OS=$(uname)
if [[ ${OS} == "Darwin" ]]; then
    brew update > /dev/null 2>&1
    brew install git cmake qt go > /dev/null 2>&1
elif [[ ${OS} == "Linux" ]]; then
    if [ -e "/etc/lsb-release" ]; then #Ubuntu
        sudo apt-get update > /dev/null 2>&1
        sudo apt-get install git cmake qt5-default go > /dev/null 2>&1
    else
        sudo yum update > /dev/null 2>&1
        sudo yum install git cmake qt go > /dev/null 2>&1
    fi
fi
echo "::set-env name=GOPATH::$(go env GOPATH)"
echo "[Complete] Install Dependencies."

#!/bin/bash
set -v

# Install Dependency
echo "Install Dependencies..."
OS=$(uname)
if [[ ${OS} == "Darwin" ]]; then
    brew update
    brew install git cmake qt go
elif [[ ${OS} == "Linux" ]]; then
    if [ -e "/etc/lsb-release" ]; then #Ubuntu
        sudo apt-get update
        sudo apt-get -y install build-essential
        sudo apt-get -y install git libssl-dev cmake qt5-default python3 language-pack-ko golang-go
    elif [ -e "/sbin/apk" ]; then #Alpine in Docker image
        apk add git cmake qt5-qtbase-dev go
    else
        sudo yum update
        sudo yum install git cmake qt golang
    fi
fi
echo "[Complete] Install Dependencies."

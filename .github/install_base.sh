#/bin/bash

# Install Dependency
OS=$(uname)
if [[ ${OS} == "Darwin" ]]; then
    brew install git cmake qt
elif [[ ${OS} == "Linux" ]]; then
    if [ -e "/etc/lsb-release" ]; then #Ubuntu
        sudo apt-get install git cmake qt5-default
    else
        sudo yum install git cmake qt
    fi
fi

# Set base path
cd `dirname ${BASH_SOURCE}`
BASE=$(pwd)
KHAIII_LIBC_DIR=$(pwd)/build/khaiii-libc

# Download Original Khaiii
cd build
git clone https://github.com/kakao/khaiii.git
cd khaiii

# Build and Install Original Khaiii
mkdir build && cd build
cmake ..
make all
cp -pfr lib ${KHAIII_LIBC_DIR}
cp -pfr ../include/khaiii ${KHAIII_LIBC_DIR}

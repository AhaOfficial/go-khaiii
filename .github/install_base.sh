#!/bin/bash
set -v

# Set base path
cd `dirname ${BASH_SOURCE}`
BASE=$(pwd)
KHAIII_LIBC_DIR=$(pwd)/build/khaiii-libc

echo "========== Set Path =========="
echo "[OS] ${OS}"
echo "[BASE] ${BASE}"
echo "[KHAIII_LIBC_DIR] ${KHAIII_LIBC_DIR}"
echo "=============================="

ls -al

# Download Original Khaiii
cd build
ls -al

git clone https://github.com/kakao/khaiii.git

cd khaiii
ls -al

# Build and Install Original Khaiii
echo "Build and Install Original Khaiii..."
mkdir build && cd build

if [ -e /etc/os-release ]; then
    if [ $(cat /etc/os-release | grep UBUNTU_CODENAME | awk -F'=' '{print $2}') = "focal" ]; then
        cmake -E env CXXFLAGS="-w" cmake ..     # Ubuntu 20.04
    else
        cmake ..
    fi
else
    cmake ..
fi

make all
make resource

ls -al lib

cp -pfr lib ${KHAIII_LIBC_DIR}
cp -pf lib/libkhaiii.* ${BASE}/..
cp -pfr share ${BASE}/..

ls -al ../include/khaiii

cp -pfr ../include/khaiii ${KHAIII_LIBC_DIR}

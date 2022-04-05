#!/bin/bash
set -v
prefix=/usr/local

# Download Original Khaiii
git clone https://github.com/kakao/khaiii.git

# Build and Install Original Khaiii
cd khaiii
echo "Build and Install Original Khaiii..."
mkdir build && cd build

OS=$(uname)
if [[ ${OS} == "Darwin" ]]; then
    cmake ..
else
    cmake -E env CXXFLAGS="-w" cmake ..
fi

make all
make resource

# Copy libraries, resources, and headers 
cp -pf lib/libkhaiii.* ${prefix}/lib
cp -pfr share/* ${prefix}/share
cp -pfr ../include/khaiii ${prefix}/include

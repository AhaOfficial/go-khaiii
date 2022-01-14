#/bin/bash

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

# Build and Install Khaiii Library for C
cd ${KHAIII_LIBC_DIR}
qmake khaiii-libc.pro
make
cp -pf khaiiic.h ${BASE}/khaiii
cp -pf libkhaiiic.* ${BASE}/khaiii
cp -pf libkhaiiic.* ${BASE}

# Run with go-khaiii
cd ${BASE}
go mod init go-khaiii
go mod tidy

# You can run this
echo "Run [ go run main.go ]"
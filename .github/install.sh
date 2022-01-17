#/bin/bash

# Set base path
cd `dirname ${BASH_SOURCE}`
BASE=$(pwd)
KHAIII_LIBC_DIR=$(pwd)/build/khaiii-libc

# Build and Install Khaiii Library for C
cd ${KHAIII_LIBC_DIR}
qmake khaiii-libc.pro
make
cp -pf khaiiic.h ${BASE}
cp -pf libkhaiiic.* ${BASE}
cp -pf khaiiic.h ${BASE}/..
cp -pf libkhaiiic.* ${BASE}/..

# Run with go-khaiii
cd ${BASE}/..
go mod init go-khaiii
go mod tidy

echo "Complete!"

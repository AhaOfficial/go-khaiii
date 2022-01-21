#/bin/bash
set -v

# Set base path
cd `dirname ${BASH_SOURCE}`
BASE=$(pwd) # .github
KHAIII_LIBC_DIR=${BASE}/build/khaiii-libc

echo "[BASE] ${BASE}"
echo "[KHAIII_LIBC_DIR] ${KHAIII_LIBC_DIR}"

# Build and Install Khaiii Library for C
cd ${KHAIII_LIBC_DIR}
qmake khaiii-libc.pro
make
cp -pf khaiiic.h ${BASE}
cp -pf libkhaiiic.* ${BASE}
cp -pf khaiiic.h ${BASE}/..
cp -pf libkhaiiic.* ${BASE}/..

ls -al ${BASE}
ls -al ${BASE}/..

echo "Build Complete!"

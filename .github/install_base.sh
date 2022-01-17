#/bin/bash

# Install Dependency
echo "Install Dependencies..."
OS=$(uname)
if [[ ${OS} == "Darwin" ]]; then
    brew install git cmake qt > /dev/null 2>&1
elif [[ ${OS} == "Linux" ]]; then
    if [ -e "/etc/lsb-release" ]; then #Ubuntu
        sudo apt-get install git cmake qt5-default > /dev/null 2>&1
    else
        sudo yum install git cmake qt > /dev/null 2>&1
    fi
fi
echo "[Complete] Install Dependencies."

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
echo "cd build"
cd build
ls -al

echo "git clone https://github.com/kakao/khaiii.git"
git clone https://github.com/kakao/khaiii.git

echo "cd khaiii"
cd khaiii
ls -al

# Build and Install Original Khaiii
echo "Build and Install Original Khaiii..."
mkdir build && cd build
cmake .. > /dev/null 2>&1
make all  > /dev/null 2>&1

echo "ls -al lib"
ls -al lib

echo "cp -pfr lib ${KHAIII_LIBC_DIR}"
cp -pfr lib ${KHAIII_LIBC_DIR}

echo "ls -al ../include/khaiii"
ls -al ../include/khaiii

echo "cp -pfr ../include/khaiii ${KHAIII_LIBC_DIR}"
cp -pfr ../include/khaiii ${KHAIII_LIBC_DIR}

#!/bin/bash
set -xeu

# symlink c and c++ compilers to build gsky
rm /usr/bin/gcc
rm /usr/bin/g++
ln -s /usr/bin/gcc-14 /usr/bin/gcc
ln -s /usr/bin/g++-14 /usr/bin/g++

#install go
wget -O go.tar.gz https://dl.google.com/go/go1.16.3.linux-amd64.tar.gz
rm -rf go && tar -xf go.tar.gz && rm -f go.tar.gz

C_INCLUDE_PATH=$(nc-config --includedir)
export C_INCLUDE_PATH
export GOROOT=/go
export GOPATH=/gsky/gopath
export PATH="$PATH:$GOROOT/bin"
mkdir -p $GOPATH

GSKY_SRC_ROOT=/gsky/gsky_src
mkdir -p $GSKY_SRC_ROOT

git clone 'https://github.com/shirei220/gsky.git' $GSKY_SRC_ROOT/gsky

(set -xeu
cd $GSKY_SRC_ROOT/gsky

./configure --prefix=/gsky --bindir=/gsky/bin --sbindir=/gsky/bin --libexecdir=/gsky/bin
make all
make install
)

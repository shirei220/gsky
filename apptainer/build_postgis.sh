#!/bin/bash
set -xeu

# symlink c and c++ compilers to build postgis
ln -s /usr/bin/gcc-11 /usr/bin/gcc
ln -s /usr/bin/g++-11 /usr/bin/g++

(set -xeu
wget https://download.osgeo.org/postgis/source/postgis-3.3.8.tar.gz
tar -xvzf postgis-3.3.8.tar.gz
cd postgis-3.3.8
./configure --with-pgconfig=/usr/lib/postgresql13/bin/pg_config
make -j4
make install
)
rm -rf postgis-3.3.8
rm -f postgis-3.3.8.tar.gz

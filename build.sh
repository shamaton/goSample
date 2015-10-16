#!/bin/sh
#set -x

#project directory
DIR=`dirname $0`
cd ${DIR}

PJ_DIR=`pwd`
INSTALL_DIR="${PJ_DIR}/src/main"

# build
export GOPATH=${PJ_DIR}

cd ${INSTALL_DIR}
go install

# recommend server setting
# export PATH=$PATH:$GOPATH/bin

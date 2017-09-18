#!/bin/bash
# NOTICE: this script can not be executed directly.
#        (1) change the svn branch;
#        (2) In step2, change the package name in *.proto files;

base="./todel"
inputdir=./protobuf/
outputdir=./proto/

baseURL=https://vcs.turbonomic.com/svn/Master/VMTurbo/
branch=$baseURL/branches/6.0.0
project="com.vmturbo.platform.sdk.common/src/main/protobuf"

set -x
# 1. get the protobuf from SVN
cd $base
svn checkout $branch/$project
rm -rf ./$inputdir/.svn

# 2. change the package name from com_dto to proto
cat $inputdir/*.proto | grep "^package" | grep "common_dt"
ret=$?
if [ $ret -eq 0 ] ; then
    echo "change the package name from com_dto to proto in the *.proto files."
    exit 1
fi

# 3. compile the protobuf to generate golang code
mkdir $outputdir
# https://github.com/golang/protobuf/issues/254
protoc --go_out=$outputdir -I=$inputdir $inputdir/*.proto


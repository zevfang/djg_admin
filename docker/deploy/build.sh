#!/usr/bin/env bash

if [ ! -n "$1" ];then
    echo "--> 警告：版本号必须填写！例如：./build.sh v1.1.3"
    exit -1
else
    echo "--> 您输入的版本号是：$1"
fi
TAG=$1

echo "--> 1.打包镜像"
docker build -t djg_admin:${TAG} .

echo "--> 2.存储镜像"
docker save -o ./image/djg_admin_${TAG}.tar djg_admin:${TAG}

echo "--> 3.删除本地镜像"
docker rmi djg_admin:${TAG}
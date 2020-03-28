#!/usr/bin/env bash

if [ ! -n "$1" ];then
    echo "--> 警告：版本号必须填写！例如：./push.sh v1.1.3"
    exit -1
else
    echo "--> 您输入的版本号是：$1"
fi
TAG=$1

scp ./image/djg_admin_${TAG}.tar root@47.92.244.167:/home/docker-ce/images
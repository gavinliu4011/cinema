#!/usr/bin/env bash
dir=`pwd`

build() {
    # 切换目录
    pushd ${dir}/$1 > /dev/null
    echo "building... ${dir}/$1"
    # 编译

    CGO_ENABLED=0  go build -o $1
    # 重置目录
    popd > /dev/null
}

build user
build api
build movie
build projection
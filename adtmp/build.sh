### 构建项目

set -ex

target=adtmp

_build() {
    go build -C cmd/server -o ../../${target} -trimpath -ldflags "-s -w" && \
        echo "构建完成!"
}

_clean() {
    if [ -f ${target} ];then
        rm -f ${target} && echo "清理完成"
    fi
}

case "$1" in
    "build"|"")
        _build
        ;;
    "clean")
        _clean
        ;;
    *)
        echo "Usage: $0 [build|clean]"
        ;;
esac
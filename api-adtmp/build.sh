### 构建项目

set -e

target=api-adtmp

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
        echo "  构建 => $0"
        echo "  清理 => $0 clean"
        ;;
esac
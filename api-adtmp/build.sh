### 构建项目

set -e


RED='\033[0;31m'
GREEN='\033[0;32m'
RESET='\033[0m'

_echo() {
    local color=$1
    shift
    echo -e "${color}$@${RESET}"
}

target=api-adtmp

_build() {
    go build -C cmd/server -o ../../${target} -trimpath -ldflags "-s -w" && \
        _echo $GREEN "构建完成 => $target !"
}

_clean() {
    if [ -f ${target} ];then
        rm -f ${target} && _echo $GREEN "$target 清理完成!"
    else
        _echo $RED "$target 已经清理!"
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

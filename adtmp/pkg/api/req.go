package api

import (
	"errors"
	"strconv"
)

var errStrToUint = errors.New("string to uint error")

// 字符串转无符号整型
func StrToUint(value string) (uint, error) {
	uintValue, err := strconv.ParseUint(value, 10, 0) // 第一个参数是字符串，第二个是进制（10 表示十进制），第三个是返回值的大小
	if err != nil {
		return 0, errStrToUint
	}
	return uint(uintValue), nil
}

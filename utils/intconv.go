package utils

import (
	"strconv"
	"strings"
)

// 字符串使用 sep 的内容进行拼接
func Int64ToString(target []int64, sep string) string {
	a := make([]string, len(target))
	for index, value := range target {
		a[index] = strconv.Itoa(int(value))
	}
	return strings.Join(a, sep)
}

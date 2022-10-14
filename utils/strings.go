package utils

import (
	"unicode/utf8"
)

// Check whether the string is too long with utf8
// 检查字符串是否过长 (utf8)
func CheckString2Long(target string, len int) bool {
	return utf8.RuneCountInString(target) >= len
}

// 检查字符串长度是否在范围内
func CheckStringInRange(target string, min, max int) bool {
	return utf8.RuneCountInString(target) <= max && utf8.RuneCountInString(target) >= min
}

func CheckString2Short(target string, len int) bool {
	return utf8.RuneCountInString(target) <= len
}

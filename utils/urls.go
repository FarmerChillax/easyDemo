package utils

import (
	"regexp"
)

// 目前仅支持 HTTP & HTTPS 协议
func CheckURL(url string) (bool, error) {
	return regexp.Match(`^(http|https)://.*?`, []byte(url))
}

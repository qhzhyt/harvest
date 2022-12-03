package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// StringFirstLower 首字母小写
func StringFirstLower(src string) string {
	return strings.ToLower(src[:1]) + src[1:]
}

// StringFirstUpper 首字母大写
func StringFirstUpper(src string) string {
	return strings.ToUpper(src[:1]) + src[1:]
}

// StringFirstIsUpper 首字母是否大写
func StringFirstIsUpper(src string) bool {
	return src[0] >= 'A' && src[0] <= 'Z'
}

// StringFirstIsLower 首字母是否小写
func StringFirstIsLower(src string) bool {
	return src[0] >= 'a' && src[0] <= 'z'
}

// StringMD5HEXString 计算字符串的 MD5 16 进制字符串
func StringMD5HEXString(str string) string {
	bytes := md5.Sum([]byte(str))
	return hex.EncodeToString(bytes[:])
}

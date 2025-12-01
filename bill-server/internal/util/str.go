package util

import "strings"

// ContainsAny 检查字符串是否包含任意一个子字符串
// str: 待检查的字符串
// arr: 子字符串列表
// 返回值: 如果字符串包含任意一个子字符串，则返回true；否则返回false
func StrContainsAny(str string, arr ...string) bool {
	for _, item := range arr {
		if strings.Contains(str, item) {
			return true
		}
	}
	return false
}

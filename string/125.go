package main

import (
	"fmt"
	"strings"
)

// 125. 验证回文串
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
// 说明：本题中，我们将空字符串定义为有效的回文串。
// https://leetcode-cn.com/problems/valid-palindrome/
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

// 使用双指针+ASCII码过滤非字母和数字的字符
// 也可以使用正则，但效率较低
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !validate(s[i]) {
			i++
		} else if !validate(s[j]) {
			j--
		} else {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
	}

	return true
}

// 是否为字母、数字
func validate(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

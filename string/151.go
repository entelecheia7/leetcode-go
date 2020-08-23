package main

import (
	"fmt"
	"strings"
)

// 151. 翻转字符串里的单词
// 给定一个字符串，逐个翻转字符串中的每个单词。
// 说明：
//     无空格字符构成一个单词。
//     输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//     如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
// https://leetcode-cn.com/problems/reverse-words-in-a-string
func main() {
	fmt.Println(reverseWords("a good   example"))
}

// 可以利用库函数，将字符串基于空格分隔为数组，反转数组，再组合为字符串

// 法一：暴力
func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	s = strings.Trim(s, " ")
	result := make([]byte, 0, len(s))
	right := len(s) // 单词的右边界，不含
	for i := len(s) - 1; i >= 0; {
		if s[i] == ' ' {
			result = append(result, []byte(s[i+1:right])...)
			result = append(result, ' ')
			i--
			for s[i] == ' ' {
				i--
			}
			right = i + 1
		} else {
			i--
		}
	}
	result = append(result, []byte(s[:right])...)
	return string(result)
}

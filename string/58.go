package main

import (
	"fmt"
)

// 58. 最后一个单词的长度
// 给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
// 如果不存在最后一个单词，请返回 0 。
// 说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。
// https://leetcode-cn.com/problems/length-of-last-word/
func main() {
	fmt.Println(lengthOfLastWord(" "))
}

func lengthOfLastWord(s string) int {
	k := len(s) - 1
	// 过滤末尾空格
	for k >= 0 && s[k] == ' ' {
		k--
	}
	end := k // 最右侧的字母位，或-1
	for k >= 0 {
		k--
		if k >= 0 && s[k] == ' ' {
			break
		}
	}
	return end - k
}

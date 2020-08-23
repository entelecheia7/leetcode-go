package main

import "fmt"

// 557. 反转字符串中的单词 III
// 给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
// 提示：
//     在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
// https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/
func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

// 暴力
// 可将字符串末尾补一个空格，以便处理最后一个单词
func reverseWords(s string) string {
	n := len(s)
	s += " "
	begin := 0 // 单词的起始位置
	result := make([]byte, 0, n)
	for i := 0; i <= n; i++ {
		if s[i] == ' ' {
			for c := i - 1; c >= begin; c-- {
				result = append(result, s[c])
			}
			result = append(result, ' ')
			begin = i + 1
		}
	}
	return string(result[:n])
}

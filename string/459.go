package main

import (
	"fmt"
	"strings"
)

// 459. 重复的子字符串
// 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
// https://leetcode-cn.com/problems/repeated-substring-pattern/
func main() {
	// fmt.Println(repeatedSubstringPattern2("abab")) // true
	fmt.Println(repeatedSubstringPattern2("aba")) // false
}

// 法一：暴力
// 如果该字符串可由子串重复n次构成，则s[0]为子串左边界，s[len(s)-1]为子串右边界
// 查找满足条件的子串，进行检查
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	subN := n / 2
	// 子串长度必定小于等于n/2，且右边界在s[:n/2]范围存在
	for i := 0; i < subN; i++ {
		if s[i] == s[n-1] && n%(i+1) == 0 {
			if validateSubString(s[:i+1], s, i+1, n) {
				return true
			}
		}
	}
	return false
}

func validateSubString(sub, s string, subLen, n int) bool {
	for i := 0; i < n; i += subLen {
		if s[i:i+subLen] != sub {
			return false
		}
	}
	return true
}

// 法二：若s非重复子串，则在字符串 (s + s)中，跳过第一个s子串查找 s，其结果必然为len(s)
func repeatedSubstringPattern2(s string) bool {
	return (strings.Index((s + s)[1:], s) + 1) != len(s)
}

// 法三：KMP
// func repeatedSubstringPattern3(s string) bool {

// 	return true
// }

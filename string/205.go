package main

import (
	"fmt"
	"strings"
)

// 205. 同构字符串
// 给定两个字符串 s 和 t，判断它们是否是同构的。
// 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
// 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
// 说明:
// 你可以假设 s 和 t 具有相同的长度。
// https://leetcode-cn.com/problems/isomorphic-strings/
func main() {
	fmt.Println(isIsomorphic2("paper", "title"))
}

// 题目要求s和t的字符相互具有唯一的映射关系
// 法一：暴力，保存s=>t、t=>s的映射关系，逐一验证
func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte) // s[i] => t[i]的映射关系
	m2 := make(map[byte]byte) // t[i] => s[i]的映射关系
	for i := 0; i < len(s); i++ {
		v, exist := m1[s[i]]
		if !exist {
			m1[s[i]] = t[i]
		} else if v != t[i] {
			return false
		}
		v, exist = m2[t[i]]
		if !exist {
			m2[t[i]] = s[i]
		} else if v != s[i] {
			return false
		}

	}
	return true
}

// 法二：对于s[i]和t[i]，它们在s和t中具有相同的下标位置。
func isIsomorphic2(s string, t string) bool {
	n := len(s)
	for i := 0; i < n; i++ {
		if strings.Index(s, s[i:i+1]) != strings.Index(t, t[i:i+1]) {
			return false
		}
	}
	return true
}

// 法三：法二的另一种写法
func isIsomorphic3(s string, t string) bool {
	n := len(s)
	var m1, m2 [256]int
	for i := 0; i < n; i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}

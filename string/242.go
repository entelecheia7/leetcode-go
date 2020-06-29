package main

import (
	"fmt"
	"strings"
)

// 242. 有效的字母异位词
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 说明:
// 你可以假设字符串只包含小写字母。
// 进阶:
// 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
// https://leetcode-cn.com/problems/valid-anagram/
func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

// 法一：使用一个额外的数组记录字符出现的频次，时间复杂度O(n)，空间复杂度O(1)
func isAnagram(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}
	m := make([]int, 26)
	for i := 0; i < ls; i++ {
		if s[i] != t[i] {
			m[s[i]-'a']++
			m[t[i]-'a']--
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// 法二：给两个字符串按同一算法排序，看生成的字符串是否相等

// 进阶：将数组换成map，map的下标可以是unicode字符

package main

import (
	"fmt"
	"strings"
)

// 290. 单词规律
// 给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
// 示例1:
// 输入: pattern = "abba", str = "dog cat cat dog"
// 输出: true
// 说明:
// 你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
// https://leetcode-cn.com/problems/word-pattern/
func main() {
	fmt.Println(wordPattern("abba", "dog cat cat fish"))
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}

// 使用哈希表保存双向映射关系，O(n)
func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	n := len(pattern)
	if len(words) != n {
		return false
	}
	p2s := make(map[byte]string, n)
	s2p := make(map[string]byte, n)
	for i := 0; i < n; i++ {
		// 检查pattern到str的映射关系
		if w, exist := p2s[pattern[i]]; !exist {
			p2s[pattern[i]] = words[i]
		} else if w != words[i] {
			return false
		}
		// 检查str到pattern的映射关系
		if p, exist := s2p[words[i]]; !exist {
			s2p[words[i]] = pattern[i]
		} else if p != pattern[i] {
			return false
		}
	}
	return true
}

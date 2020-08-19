package main

import "fmt"

// 387. 字符串中的第一个唯一字符
// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
// 提示：你可以假定该字符串只包含小写字母。
// https://leetcode-cn.com/problems/first-unique-character-in-a-string/
func main() {
	fmt.Println(firstUniqChar("loveleetcode")) // 2
}

// 循环+数组
// 只有小写字母可以用数组，字符范围较大可以用map
func firstUniqChar(s string) int {
	n := len(s)
	if s == "" {
		return -1
	}
	m := make([]int, 26)
	for i := 0; i < n; i++ {
		m[s[i]-'a']++
	}
	unique := -1
	for i := 0; i < n; i++ {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}
	return unique
}

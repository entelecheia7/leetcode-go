package main

import "fmt"

// 438. 找到字符串中所有字母异位词
// 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
// 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
// 说明：
//     字母异位词指字母相同，但排列不同的字符串。
//     不考虑答案输出的顺序。
// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc")) // [0 6]
	fmt.Println(findAnagrams("abab", "ab"))        // [0 1 2]
	fmt.Println(findAnagrams("baa", "aa"))         // [1]
}

// 法一：滑动窗口
func findAnagrams(s string, p string) (result []int) {
	sl, pl := len(s), len(p)
	if pl > sl {
		return nil
	}
	need := [26]int{}
	for i := 0; i < pl; i++ {
		need[p[i]-'a']++
	}
	validType := 0
	for i := 0; i < 26; i++ {
		if need[i] > 0 {
			validType++
		}
	}
	window := [26]int{}
	valid := 0
	for i := 0; i < sl; i++ {
		// 更新窗口右侧新边界字母
		w := s[i] - 'a'
		if need[w] > 0 {
			window[w]++
			if window[w] == need[w] {
				valid++
			}
		}
		// 收缩窗口
		if i >= pl-1 {
			if valid == validType {
				result = append(result, i+1-pl)
			}
			left := s[i+1-pl] - 'a'
			if need[left] > 0 {
				if window[left] == need[left] {
					valid--
				}
				window[left]--
			}
		}
	}
	return result
}

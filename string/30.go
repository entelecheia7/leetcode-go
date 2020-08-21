package main

import "fmt"

// 30. 串联所有单词的子串
// 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/
func main() {
	// fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"})) // [0 9]
	// fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"})) // [6 9 12]
	fmt.Println(findSubstring2("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})) // [8]
	fmt.Println(findSubstring4("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"})) // [8]
}

// 法一：暴力
// 将问题转化为一个固定长度的字符串是否可以由一个字符串数组的元素组成
// 循环遍历s，对每一个以s[i]起始的，长度为 len(words)*len(words[0])的字符串进行检查
func findSubstring(s string, words []string) (result []int) {
	if len(words) == 0 {
		return nil
	}
	wordLen := len(words[0])
	matchLen := wordLen * len(words)
	wordsLength := len(words)
	if len(s) < matchLen {
		return nil
	}
	wordsMap := make(map[string]int) // 统计wrods中的单次和数量
	for _, w := range words {
		wordsMap[w]++
	}
	n := len(s)
	for i := 0; i <= n-matchLen; i++ {
		start := i
		if wordsMap[s[start:start+wordLen]] > 0 {
			wm := make(map[string]int, len(wordsMap))
			matched := 0
			for start+wordLen*(wordsLength-matched) <= n && matched < wordsLength {
				sub := s[start : start+wordLen]
				if freq, exist := wordsMap[sub]; exist && wm[sub] < freq {
					wm[sub]++
					matched++
					start += wordLen
				} else {
					break
				}
			}
			if matched == wordsLength {
				result = append(result, i)
			}
		}
	}
	return result
}

// 法二：理解为一个滑动滑动窗口
// 滑动窗口的大小是 len(words)*len(words[0])
// 设words中单次的长度是 m，以 s[0:m)中的元素作为滑动窗口的左边界，每次滑动单词长度 m
// 得到一个完全匹配的子串时，滑动 m 长度
// 匹配时需要倒序，一旦遇到不匹配的子串，可以直接跳到这个子串之后
// 当不匹配时，滑动的长度是 (len(words)-matched)*len(words[0])
func findSubstring2(s string, words []string) (result []int) {
	if len(words) == 0 {
		return nil
	}
	wordLen := len(words[0])
	matchLen := wordLen * len(words)
	wordsLength := len(words)
	if len(s) < matchLen {
		return nil
	}
	wordsMap := make(map[string]int) // 统计wrods中的单次和数量
	for _, w := range words {
		wordsMap[w]++
	}
	n := len(s)
	for i := 0; i < wordLen; i++ {
		start := i // 滑动窗口的起始位置
		for start+matchLen <= n {
			unmatched := wordsLength
			wm := make(map[string]int, len(wordsMap))
			// 倒序判断，一旦遇到不匹配的子串，可以直接跳到这个子串之后
			k := start + matchLen // key表示在当前窗口中子串匹配到的位置
			for unmatched > 0 {
				sub := s[k-wordLen : k]
				wm[sub]++
				if wm[sub] > wordsMap[sub] {
					break
				}
				unmatched--
				k -= wordLen
			}
			moveStep := 1
			if unmatched == 0 {
				result = append(result, start)
			} else {
				moveStep = getMax(1, unmatched)
			}
			start += moveStep * wordLen
		}
	}
	return result
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

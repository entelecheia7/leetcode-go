package main

import (
	"fmt"
)

// 76. 最小覆盖子串
// 给你一个字符串 S、一个字符串 T 。请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。
// 提示：
//     如果 S 中不存这样的子串，则返回空字符串 ""。
//     如果 S 中存在这样的子串，我们保证它是唯一的答案。
// https://leetcode-cn.com/problems/minimum-window-substring/
func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC")) // BANC
}

// 法一：暴力搜索，时间复杂度高，O(m*n)

// 法二：滑动窗口
// 先确定窗口大小：固定左侧，右侧滑动，当s包含t的所有元素时，就确定了初始窗口大小
// 当窗口大小确定时，开始记录窗口大小并试图移动左侧窗口
func minWindow(s string, t string) (result string) {
	ls, lt := len(s), len(t)
	if ls < lt || lt == 0 || ls == 0 {
		return
	}
	var need, window [128]int // map是一种更通用的做法，但数组的效率更高
	for i := 0; i < lt; i++ {
		need[t[i]]++ // 记录需要的元素及数量
	}
	needType := 0 // 记录需要的元素种类
	for i := 0; i < 128; i++ {
		if need[i] > 0 {
			needType++
		}
	}
	valid := 0
	left, right := 0, 0
	for right < ls {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 窗口包含所有所需元素时，进行记录，并移动左侧边界
		for valid == needType {
			// 计算长度
			if result == "" || right-left < len(result) {
				result = s[left:right]
			}
			// 滑动左侧边界
			c = s[left]
			if need[c] > 0 {
				if need[c] == window[c] {
					valid--
				}
				window[s[left]]--
			}
			left++
		}
	}

	return
}

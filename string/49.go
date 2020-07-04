package main

import (
	"fmt"
	"sort"
)

// 49. 字母异位词分组
// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
// 示例:
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
// 说明：
//     所有输入均为小写字母。
//     不考虑答案输出的顺序。
// https://leetcode-cn.com/problems/group-anagrams/
func main() {
	fmt.Println(groupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// 法一：字符串排序后的值作为map的键，进行统计
// 时间O(n*k*logk)，k是最长的字符串长度
// 空间O(nk)
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	m := make(map[string][]string)
	for _, str := range strs {
		byt := SortByte(str)
		sort.Sort(byt)
		m[byt.String()] = append(m[byt.String()], str)
	}
	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

type SortByte []byte

func (s SortByte) Less(i, j int) bool { return s[i] < s[j] }
func (s SortByte) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortByte) Len() int           { return len(s) }
func (s SortByte) String() string     { return string(s) }

// 法二：统计每个字符串的频率，进行比较
// O(nk)
func groupAnagrams2(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	m := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for i := 0; i < len(str); i++ {
			count[str[i]-'a']++
		}
		m[count] = append(m[count], str)
	}
	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

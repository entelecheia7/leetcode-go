package main

import "fmt"

// 17. 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 示例:
// 输入："23"
// 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// 说明:
// 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations2("23"))
}

var digitsRelation = [][]string{
	0: nil,
	1: nil,
	2: []string{"a", "b", "c"},
	3: []string{"d", "e", "f"},
	4: []string{"g", "h", "i"},
	5: []string{"j", "k", "l"},
	6: []string{"m", "n", "o"},
	7: []string{"p", "q", "r", "s"},
	8: []string{"t", "u", "v"},
	9: []string{"w", "x", "y", "z"},
}

// 法一：digits中的每一个数字都是在上一步的结果上添加对应的字母
// 通过循环得出所有结果，空间复杂度较大
func letterCombinations(digits string) (result []string) {
	n := len(digits)
	if n == 0 {
		return nil
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			result = digitsRelation[digits[i]-'0']
		} else {
			tmp := make([]string, 0, len(result)*len(digitsRelation[digits[i]-'0']))
			for _, r := range result {
				for _, s := range digitsRelation[digits[i]-'0'] {
					tmp = append(tmp, r+s)
				}
			}
			result = tmp
		}
	}
	return result
}

// 法二：回溯
// best
func letterCombinations2(digits string) (result []string) {
	n := len(digits)
	if n == 0 {
		return nil
	}
	letterCombinationsHelper(digits, "", &result)
	return result
}

func letterCombinationsHelper(digits string, cur string, result *[]string) {
	if digits == "" {
		*result = append(*result, cur)
		return
	}
	for _, s := range digitsRelation[digits[0]-'0'] {
		letterCombinationsHelper(digits[1:], cur+s, result)
	}
}

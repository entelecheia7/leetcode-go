package main

import (
	"fmt"
)

// 299. 猜数字游戏
// 你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：
// 你写出一个秘密数字，并请朋友猜这个数字是多少。
// 朋友每猜测一次，你就会给他一个提示，告诉他的猜测数字中有多少位属于数字和确切位置都猜对了（称为“Bulls”, 公牛），有多少位属于数字猜对了但是位置不对（称为“Cows”, 奶牛）。
// 朋友根据提示继续猜，直到猜出秘密数字。
// 请写出一个根据秘密数字和朋友的猜测数返回提示的函数，返回字符串的格式为 xAyB ，x 和 y 都是数字，A 表示公牛，用 B 表示奶牛。
// xA 表示有 x 位数字出现在秘密数字中，且位置都与秘密数字一致。
// yB 表示有 y 位数字出现在秘密数字中，但位置与秘密数字不一致。
// 请注意秘密数字和朋友的猜测数都可能含有重复数字，每位数字只能统计一次。
// 说明: 你可以假设秘密数字和朋友的猜测数都只包含数字，并且它们的长度永远相等。
// https://leetcode-cn.com/problems/bulls-and-cows
func main() {
	fmt.Println(getHint2("1123", "0111"))
}

// 法一：两次遍历+数组
// 遍历guess和secret，记录Bulls，对于miss的数字使用两个数组保存，再一次遍历统计cows
func getHint(secret string, guess string) string {
	if secret == "" || guess == "" {
		return "0A0B"
	}
	bulls, cows := 0, 0
	length := len(guess)
	missedS, missedG := make([]int, 10), make([]int, 10)
	for i := 0; i < length; i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			missedG[guess[i]-'0']++
			missedS[secret[i]-'0']++
		}
	}
	for c, n := range missedS {
		if missedG[c] > 0 {
			cows += getMin(n, missedG[c])
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 法二：一次遍历+数组
// 关键在于使用一个数组记录字符在secret和guess中出现的频率，guess字符+1，secret字符-1
func getHint2(secret string, guess string) string {
	if secret == "" || guess == "" {
		return "0A0B"
	}
	bulls, cows := 0, 0
	length := len(guess)
	numbers := make([]int, 10)
	for i := 0; i < length; i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			s, g := secret[i]-'0', guess[i]-'0'
			if numbers[s] > 0 {
				cows++
			}
			if numbers[g] < 0 {
				cows++
			}
			numbers[g]++
			numbers[s]--
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

package main

import (
	"fmt"
	"math"
)

// 818. 赛车
// 你的赛车起始停留在位置 0，速度为 +1，正行驶在一个无限长的数轴上。（车也可以向负数方向行驶。）
// 你的车会根据一系列由 A（加速）和 R（倒车）组成的指令进行自动驾驶 。
// 当车得到指令 "A" 时, 将会做出以下操作： position += speed, speed *= 2。
// 当车得到指令 "R" 时, 将会做出以下操作：如果当前速度是正数，则将车速调整为 speed = -1 ；否则将车速调整为 speed = 1。  (当前所处位置不变。)
// 例如，当得到一系列指令 "AAR" 后, 你的车将会走过位置 0->1->3->3，并且速度变化为 1->2->4->-1。
// 现在给定一个目标位置，请给出能够到达目标位置的最短指令列表的长度。
// 说明:
//     1 <= target（目标位置） <= 10000。
// https://leetcode-cn.com/problems/race-car/
func main() {
	fmt.Println(racecar(3)) // 2
	fmt.Println(racecar(6)) // 5
}

// 法一：动态规划
// dp[i]表示target是i时的最短指令长度
// target = 2^n-1时，指令均为A，这时一定是最短指令
// 当target是其他位置时，假设2^(k-1) <= target < 2^k
// 那么可以先走(k-1)次A，到达位置(2^(k-1)-1)，通过R转向，再走m个A，再走R转向，递归走剩余路程
// 也可以走k次A，到达(2^k-1)，通过R转向，再走剩余路程:(i<<k)-1-target
var dp [10001]int

func racecar(target int) int {
	if dp[target] > 0 {
		return dp[target]
	}
	k := int(math.Floor(math.Log2(float64(target)))) + 1
	if target+1 == (1 << k) {
		return k
	}
	// 走k次A，到达(2^k-1)，通过R转向，再走剩余路程
	dp[target] = k + 1 + racecar((1<<k)-1-target)
	// 走(k-1)次A，到达位置(2^(k-1)-1)，通过R转向，再走m个A，再走R转向，递归走剩余路程
	// m的取值范围是[0, k-1)
	for m := 0; m < k-1; m++ {
		dp[target] = getMin(dp[target], k+m+1+racecar(target-(1<<(k-1))+(1<<m)))
	}
	return dp[target]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

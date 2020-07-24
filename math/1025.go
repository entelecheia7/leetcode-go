package main

import "fmt"

// 1025. 除数博弈
// 爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。
// 最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：
//     选出任一 x，满足 0 < x < N 且 N % x == 0 。
//     用 N - x 替换黑板上的数字 N 。
// 如果玩家无法执行这些操作，就会输掉游戏。
// 只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。假设两个玩家都以最佳状态参与游戏。
// 提示：
//     1 <= N <= 1000
// https://leetcode-cn.com/problems/divisor-game/
func main() {

}

// 数学归纳法！
// 数字N如果是奇数，它的约数必然都是奇数；若为偶数，则其约数可奇可偶。
// 无论N初始为多大的值，游戏最终只会进行到N=2时结束，那么谁轮到N=2时谁就会赢。
// 因为爱丽丝先手，N初始若为偶数，爱丽丝则只需一直选1，使鲍勃一直面临N为奇数的情况，这样爱丽丝稳赢；
// N初始若为奇数，那么爱丽丝第一次选完之后N必为偶数，那么鲍勃只需一直选1就会稳赢。
// https://leetcode-cn.com/problems/divisor-game/solution/qi-shi-shi-yi-dao-shu-xue-ti-by-coder233/
func divisorGame2(N int) bool {
	return N%2 == 0
}

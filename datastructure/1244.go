package main

import "fmt"

// 1244. 力扣排行榜
// 新一轮的「力扣杯」编程大赛即将启动，为了动态显示参赛者的得分数据，需要设计一个排行榜 Leaderboard。
// 请你帮忙来设计这个 Leaderboard 类，使得它有如下 3 个函数：
//     addScore(playerId, score)：
//         假如参赛者已经在排行榜上，就给他的当前得分增加 score 点分值并更新排行。
//         假如该参赛者不在排行榜上，就把他添加到榜单上，并且将分数设置为 score。
//     top(K)：返回前 K 名参赛者的 得分总和。
//     reset(playerId)：将指定参赛者的成绩清零。题目保证在调用此函数前，该参赛者已有成绩，并且在榜单上。
// 请注意，在初始状态下，排行榜是空的。
// 提示：
//     1 <= playerId, K <= 10000
//     题目保证 K 小于或等于当前参赛者的数量
//     1 <= score <= 100
//     最多进行 1000 次函数调用
// https://leetcode-cn.com/problems/design-a-leaderboard
func main() {
	board := Constructor()
	board.AddScore(1, 73)
	board.AddScore(2, 56)
	board.AddScore(3, 39)
	board.AddScore(4, 51)
	board.AddScore(5, 4)
	fmt.Println(board.Top(1))
	board.Reset(1)
	board.Reset(2)
	board.AddScore(2, 51)
	fmt.Println(board.Top(3))
}

type Leaderboard struct {
	players map[int]int // playerId =>score
	scores  []int       // 排好序的score
}

func Constructor() Leaderboard {
	return Leaderboard{
		players: make(map[int]int),
	}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
	if _, exist := this.players[playerId]; exist {
		oldScore := this.players[playerId]
		this.players[playerId] += score
		this.deleteScore(oldScore)
		this.insertScore(this.players[playerId])
	} else {
		this.players[playerId] = score
		this.insertScore(score)
	}
}

func (this *Leaderboard) Top(K int) (scores int) {
	i := len(this.scores) - 1
	for K > 0 {
		scores += this.scores[i]
		i--
		K--
	}
	return scores
}

func (this *Leaderboard) Reset(playerId int) {
	this.deleteScore(this.players[playerId])
	delete(this.players, playerId)
}

// score总是存在
func (this *Leaderboard) deleteScore(score int) {
	i := this.search(score)
	this.scores = append(this.scores[:i], this.scores[i+1:]...)
}
func (this Leaderboard) search(score int) (index int) {
	index = -1
	left, right := 0, len(this.scores)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if this.scores[mid] == score {
			return mid
		} else if this.scores[mid] < score {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}
func (this *Leaderboard) insertScore(score int) {
	n := len(this.scores)
	if n == 0 || this.scores[n-1] <= score {
		this.scores = append(this.scores, score)
		return
	} else if this.scores[0] >= score {
		this.scores = append([]int{score}, this.scores...)
		return
	}
	// 找到小于等于score的最大元素
	left, right := 0, n-1
	target := -1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if this.scores[mid] <= score {
			target = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	tmp := make([]int, n+1)
	copy(tmp, this.scores[:target+1])
	tmp[target+1] = score
	copy(tmp[target+2:], this.scores[target+1:])
	this.scores = tmp
}

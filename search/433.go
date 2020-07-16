package main

import (
	"fmt"
	"math"
)

// 433. 最小基因变化
// 一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
// 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
// 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
// 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
// 现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。
// 注意:
//     起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
//     所有的目标基因序列必须是合法的。
//     假定起始基因序列与目标基因序列是不一样的。
// 	https://leetcode-cn.com/problems/minimum-genetic-mutation/
func main() {
	fmt.Println(minMutation2("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}))                                     // 1
	fmt.Println(minMutation2("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))             // 2
	fmt.Println(minMutation2("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}))             // 3
	fmt.Println(minMutation2("AACCGGTT", "AAACGGTA", []string{"AACCGATT", "AACCGATA", "AAACGATA", "AAACGGTA"})) // 4
}

// 思路：bank中存在的和start及中间结果只差1位的序列，就是下一步的可转换序列
// 法一：DFS。通过递归进行尝试
func minMutation(start string, end string, bank []string) (minMu int) {
	bankMap := make(map[string]bool, len(bank))
	for _, s := range bank {
		bankMap[s] = true
	}
	if _, exist := bankMap[end]; !exist {
		return -1
	}
	visited := make(map[string]bool, len(bank)) // 避免产生转换环路
	minMu = math.MaxInt64
	minMutationDFSHelper(start, end, []string{"A", "C", "G", "T"}, bankMap, 0, visited, &minMu)
	if minMu == math.MaxInt64 {
		return -1
	}
	return minMu
}

func minMutationDFSHelper(start, end string, material []string, bankMap map[string]bool, changeNum int, visited map[string]bool, minMu *int) {
	if changeNum >= *minMu { // 剪枝，提前结束不可能最短的路径
		return
	}
	if start == end { // 转换成功
		*minMu = getMin(*minMu, changeNum)
		return
	}

	for i := 0; i < 8; i++ {
		// 尝试寻找转换的下一个序列
		for _, m := range material {
			if start[i:i+1] != m {
				next := start[:i] + m + start[i+1:]
				if _, exist := bankMap[next]; exist && !visited[next] {
					visited[next] = true
					minMutationDFSHelper(next, end, material, bankMap, changeNum+1, visited, minMu)
					visited[next] = false
				}
			}
		}
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 法二：BFS
// 将从start开始，转换经过的序列保存至队列
func minMutation2(start string, end string, bank []string) (changeNum int) {
	bankMap := make(map[string]bool, len(bank))
	for _, s := range bank {
		bankMap[s] = true
	}
	if _, exist := bankMap[end]; !exist {
		return -1
	}
	visited := make(map[string]bool, len(bank)) // 避免产生转换环路
	visited[start] = true
	level := 0
	queue := []string{start}
	material := []string{"A", "C", "G", "T"}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if queue[i] == end {
				return level
			}
			for j := 0; j < 8; j++ {
				for _, m := range material {
					if queue[i][j:j+1] != m {
						next := queue[i][:j] + m + queue[i][j+1:]
						if _, exist := bankMap[next]; exist && !visited[next] {
							visited[next] = true
							queue = append(queue, next)
						}
					}
				}
			}
		}
		queue = queue[size:]
		level++
	}
	return -1
}

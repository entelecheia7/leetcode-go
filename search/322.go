package main

import (
	"fmt"
	"sort"
)

// 332. 重新安排行程
// 给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。
// 说明:
//     如果存在多种有效的行程，你可以按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
//     所有的机场都用三个大写字母表示（机场代码）。
//     假定所有机票至少存在一种合理的行程。
// https://leetcode-cn.com/problems/reconstruct-itinerary/
func main() {
	fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))                                                                                                 // ["JFK", "MUC", "LHR", "SFO", "SJC"]
	fmt.Println(findItinerary([][]string{{"JFK", "ATL"}, {"ORD", "PHL"}, {"JFK", "ORD"}, {"PHX", "LAX"}, {"LAX", "JFK"}, {"PHL", "ATL"}, {"ATL", "PHX"}}))                                                 // ["JFK","ATL","PHX","LAX","JFK","ORD","PHL","ATL"]
	fmt.Println(findItinerary([][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}})) // ["JFK","ANU","EZE","AXA","TIA","ANU","JFK","TIA","ANU","TIA","JFK"]
	fmt.Println(findItinerary([][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}))
}

// 法一：DFS
// 由于必然存在一条有效路径，所以算法不需要回溯（加入到结果集里的元素不需要删除）
func findItinerary(tickets [][]string) (path []string) {
	n := len(tickets)
	// 构建有向图
	graph := make(map[string][]string)
	for i := 0; i < n; i++ {
		t := tickets[i]
		graph[t[0]] = append(graph[t[0]], t[1])
	}
	// 对graph的value进行排序
	for k, v := range graph {
		if len(v) == 1 {
			continue
		}
		sort.Strings(graph[k])
	}
	var dfsHelper func(cur string)
	dfsHelper = func(cur string) {
		for len(graph[cur]) > 0 {
			next := graph[cur][0]
			graph[cur] = graph[cur][1:]
			dfsHelper(next)
		}
		// 这里无法通过递增k来组织path，因为同一层级的元素k是相同的，会被覆盖
		path = append(path, cur)
	}

	dfsHelper("JFK")
	// 反转path
	for i := 0; i < len(path)/2; i++ {
		path[i], path[n-i] = path[n-i], path[i]
	}
	return path
}

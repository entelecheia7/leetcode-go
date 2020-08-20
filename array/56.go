package main

import "fmt"

// 56. 合并区间
// 给出一个区间的集合，请合并所有重叠的区间。
// 提示：
//     intervals[i][0] <= intervals[i][1]
// https://leetcode-cn.com/problems/merge-intervals/
func main() {
	// fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	// fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}))
}

// 法一：将输入按照元素的第一个元素排序
// 快排
func merge(intervals [][]int) (merged [][]int) {
	if len(intervals) <= 1 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1)
	i := 0
	merged = make([][]int, 0, len(intervals))
	for i < len(intervals) {
		start := intervals[i][0]
		end := intervals[i][1]
		// fmt.Println("===", i, start, end)
		for i < len(intervals)-1 && end >= intervals[i+1][0] {
			i++
			start = getMin(start, intervals[i][0])
			end = getMax(end, intervals[i][1])
		}
		merged = append(merged, []int{start, end})
		i++
	}
	return merged
}
func quickSort(intervals [][]int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(intervals, left, right)
	quickSort(intervals, left, pivot-1)
	quickSort(intervals, pivot+1, right)
}
func partition(intervals [][]int, left, right int) (pivot int) {
	standard := intervals[right]
	pos := left // 小于standard元素要放置的地方
	for i := left; i < right; i++ {
		if intervals[i][0] < standard[0] {
			intervals[pos], intervals[i] = intervals[i], intervals[pos]
			pos++
		}
	}
	intervals[pos], intervals[right] = intervals[right], intervals[pos]
	return pos
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

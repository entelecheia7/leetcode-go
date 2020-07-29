package main

import (
	"fmt"
	"sort"
)

// 621. 任务调度器
// 给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。
// 然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
// 你需要计算完成所有任务所需要的最短时间。
// 提示：
//     任务的总个数为 [1, 10000]。
//     n 的取值范围为 [0, 100]。
// https://leetcode-cn.com/problems/task-scheduler/
func main() {
	// fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)) // 8
	// fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0)) // 6
	// fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 1)) // 12
	fmt.Println(leastInterval([]byte{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D', 'E', 'E', 'F', 'F', 'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J', 'K', 'K', 'L', 'L', 'M', 'M', 'N', 'N', 'O', 'O', 'P', 'P', 'Q', 'Q', 'R', 'R', 'S', 'S', 'T', 'T', 'U', 'U', 'V', 'V', 'W', 'W', 'X', 'X', 'Y', 'Y', 'Z', 'Z'}, 2)) // 52
}

// 法一：贪心+排序
// 任务调度器的最短时间依赖于于数量最多的那个任务c，
// 将任务按照数量排序，将时间片以 n+1 的长度划分，每个时间片只能执行一种任务一次
func leastInterval(tasks []byte, n int) (time int) {
	if len(tasks) <= 1 || n == 0 {
		return len(tasks)
	}
	var task taskList
	count := 0
	for _, t := range tasks {
		task[t-'A'].task = t
		task[t-'A'].num++
		count++
	}
	// 按照任务数量从大到小排序
	sort.Sort(&task)
	// 完成数量最多的任务至少需要的时间
	time = (task[0].num-1)*(n+1) + 1
	// 如果有任务数量和最多数量相等的任务，那就需要+1
	i := 1
	for i < 26 && task[i].num == task[0].num {
		time++
		i++
	}
	// 如果计算结果小于数组数量，说明任务的种类和间隔时间要求对任务调度没有影响。
	if time < len(tasks) {
		return len(tasks)
	}
	return time
}

type taskList [26]taskData
type taskData struct {
	task byte
	num  int
}

func (t taskList) Len() int {
	return 26
}
func (t taskList) Less(i, j int) bool {
	return t[i].num > t[j].num
}
func (t *taskList) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

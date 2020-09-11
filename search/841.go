package main

import "fmt"

// 841. 钥匙和房间
// 有 N 个房间，开始时你位于 0 号房间。每个房间有不同的号码：0，1，2，...，N-1，并且房间里可能有一些钥匙能使你进入下一个房间。
// 在形式上，对于每个房间 i 都有一个钥匙列表 rooms[i]，每个钥匙 rooms[i][j] 由 [0,1，...，N-1] 中的一个整数表示，其中 N = rooms.length。 钥匙 rooms[i][j] = v 可以打开编号为 v 的房间。
// 最初，除 0 号房间外的其余所有房间都被锁住。
// 你可以自由地在房间之间来回走动。
// 如果能进入每个房间返回 true，否则返回 false。
// 提示：
//     1 <= rooms.length <= 1000
//     0 <= rooms[i].length <= 1000
//     所有房间中的钥匙数量总计不超过 3000。
// https://leetcode-cn.com/problems/keys-and-rooms/
func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))           // true
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}})) // false
}

// 法一：DFS
func canVisitAllRooms(rooms [][]int) (result bool) {
	visited := map[int]bool{0: true}
	n := len(rooms)
	dfsHelper(rooms, 0, n, visited, &result)
	return result
}

// cur表示当前所在房间，n表示房间的总数
func dfsHelper(rooms [][]int, cur, n int, visited map[int]bool, result *bool) {
	visited[cur] = true
	if len(visited) == n {
		*result = true
		return
	}
	nextRooms := rooms[cur]
	for i := 0; i < len(nextRooms); i++ {
		if !visited[nextRooms[i]] {
			dfsHelper(rooms, nextRooms[i], n, visited, result)
		}
	}
}

// 法二：BFS
func canVisitAllRooms2(rooms [][]int) bool {
	visited := map[int]bool{0: true}
	n := len(rooms)
	queue := rooms[0]
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if !visited[cur] {
				queue = append(queue, rooms[cur]...)
				visited[cur] = true
			}
		}
		queue = queue[size:]
	}
	return len(visited) == n
}

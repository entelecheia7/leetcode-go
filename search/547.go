package main

import "fmt"

// 547. 朋友圈
// 班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。
// 给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数.
// 注意：
// 		N 在[1,200]的范围内。
// 		对于所有学生，有M[i][i] = 1。
// 		如果有M[i][j] = 1，则有M[j][i] = 1。
// https://leetcode-cn.com/problems/friend-circles/
func main() {
	fmt.Println(findCircleNum2([][]int{
		{1, 1, 0},
		{1, 1, 1},
		{0, 1, 1},
	})) // 1
}

// 法一：dfs
func findCircleNum(M [][]int) (count int) {
	if len(M) == 0 {
		return 0
	}
	n := len(M)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			findCircleNumDFSHelper(M, visited, i, n)
			count++
		}
	}
	return count
}
func findCircleNumDFSHelper(M [][]int, visited []bool, i, n int) {
	for j := 0; j < n; j++ {
		if M[i][j] == 1 && !visited[j] {
			visited[j] = true
			findCircleNumDFSHelper(M, visited, j, n)
		}
	}
}

// 法二：并查集
func findCircleNum2(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	n := len(M)
	uf := NewUnionFind(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	return uf.getCount()
}

type unionFind struct {
	parent []int
	count  int
}

func NewUnionFind(n int) *unionFind {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}

	return &unionFind{
		parent: p,
		count:  n,
	}
}

// 返回集合的根元素
func (this unionFind) find(p int) int {
	root := p
	for root != this.parent[root] {
		root = this.parent[root]
	}
	// 压缩路径
	for p != this.parent[p] {
		next := this.parent[p]
		this.parent[p] = root
		p = next
	}
	return root
}
func (this *unionFind) union(x, y int) {
	rootX := this.find(x)
	rootY := this.find(y)
	if rootX == rootY {
		return
	}
	this.parent[rootX] = rootY
	this.count--
}
func (this unionFind) getCount() int {
	return this.count
}

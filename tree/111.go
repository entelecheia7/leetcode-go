package main

import (
	"fmt"
)

// 111. 二叉树的最小深度
// 给定一个二叉树，找出其最小深度。
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
// 说明: 叶子节点是指没有子节点的节点。
// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(minDepth(tree))
	fmt.Println(minDepth2(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归，O(n)
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return getMin(minDepth(root.Left), minDepth(root.Right)) + 1
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 法二：广度优先搜索，最差情况下才是O(n)
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	curDepth := 0
	for len(queue) > 0 {
		curDepth++
		levelNum := len(queue)
		for i := 0; i < levelNum; i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				return curDepth
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[levelNum:]
	}
	return curDepth
}

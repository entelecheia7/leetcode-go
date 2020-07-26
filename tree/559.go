package main

import "fmt"

// 559. N叉树的最大深度
// 给定一个 N 叉树，找到其最大深度。
// 最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
// 说明:
//     树的深度不会超过 1000。
//     树的节点总不会超过 5000。
// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/
func main() {
	tree := &Node{
		Val: 1,
		Children: []*Node{
			{3, []*Node{
				{5, nil},
				{6, nil},
			}},
			{2, nil},
			{4, nil},
		},
	}
	fmt.Println(maxDepth(tree))
}

// 法一：递归，best
func maxDepth(root *Node) (maxDep int) {
	if root == nil {
		return 0
	}
	for _, child := range root.Children {
		maxDep = getMax(maxDep, maxDepth(child))
	}
	return maxDep + 1
}

// 法二：层序遍历
func maxDepth2(root *Node) int {
	if root == nil {
		return 0
	}
	queue := []*Node{root}
	depth := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			for _, child := range cur.Children {
				if child != nil {
					queue = append(queue, child)
				}
			}
		}
		queue = queue[size:]
		depth++
	}
	return depth
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	Val      int
	Children []*Node
}

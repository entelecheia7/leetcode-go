package main

import (
	"fmt"
)

// 429. N叉树的层序遍历
// 给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
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
	fmt.Println(levelOrder(tree))
	fmt.Println(levelOrder2(tree))
}

type Node struct {
	Val      int
	Children []*Node
}

// 法一：利用队列，循环
func levelOrder(root *Node) (result [][]int) {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) != 0 {
		n := len(queue) // 本层节点数量
		val := []int{}  // 本层节点值
		for i := 0; i < n; i++ {
			val = append(val, queue[i].Val)
			if queue[i].Children != nil {
				queue = append(queue, queue[i].Children...)
			}
		}
		queue = queue[n:]
		result = append(result, val)
	}

	return
}

// 法二：递归，增加level参数表示层级
func levelOrder2(root *Node) (result [][]int) {
	if root == nil {
		return nil
	}
	helper(root, 0, &result)
	return
}
func helper(root *Node, level int, result *[][]int) {
	if len(*result) <= level {
		*result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], root.Val)
	for _, child := range root.Children {
		helper(child, level+1, result)
	}
}

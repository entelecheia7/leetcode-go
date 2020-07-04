package main

import (
	"fmt"
)

// 589. N叉树的前序遍历
// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
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

	fmt.Println(preorder(tree))
	fmt.Println(preorder2(tree))
}

type Node struct {
	Val      int
	Children []*Node
}

// 法一：递归
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	for _, child := range root.Children {
		result = append(result, preorder(child)...)
	}
	return result
}

// 法二：栈+循环
func preorder2(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	result := []int{}
	for len(stack) > 0 {
		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, back.Val)
		for i := len(back.Children) - 1; i >= 0; i-- {
			stack = append(stack, back.Children[i])
		}
	}

	return result
}

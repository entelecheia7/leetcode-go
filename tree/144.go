package main

import (
	"fmt"
)

// 144. 二叉树的前序遍历
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
func main() {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(preorderTraversal(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：利用栈，双百
func preorderTraversal(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// 法二：递归，常规方法
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)
	return result
}

type colorNode struct {
	node  *TreeNode
	color int
}

// 法三：颜色遍历
func preorderTraversal3(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	white, gray := 0, 1
	stack := []colorNode{{root, white}}
	for len(stack) > 0 {
		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if back.node != nil {
			if back.color == white {
				stack = append(stack, colorNode{back.node.Right, white})
				stack = append(stack, colorNode{back.node.Left, white})
				back.color = gray
				stack = append(stack, back)
			} else {
				result = append(result, back.node.Val)
			}
		}
	}
	return
}

package main

import (
	"fmt"
)

// 145. 二叉树的后序遍历
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
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

	fmt.Println(postorderTraversal(tree))
	fmt.Println(postorderTraversal2(tree))
	fmt.Println(postorderTraversal3(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := postorderTraversal(root.Left)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

// 法二：迭代
// 由于后序遍历的顺序是：左-右-根，所以将根节点压栈两次
// 第一个根节点用于将右、左节点压栈，第二个用于输出根节点的值
func postorderTraversal2(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root, root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) > 0 && cur == stack[len(stack)-1] {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
				stack = append(stack, cur.Left)
			}
		} else {
			result = append(result, cur.Val)
		}
	}
	return
}

// 法三：颜色遍历
type colorNode struct {
	node  *TreeNode
	color int
}

func postorderTraversal3(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	white, gray := 0, 1
	stack := []colorNode{colorNode{root, white}}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.node != nil {
			if cur.color == white {
				cur.color = gray
				stack = append(stack, cur)
				stack = append(stack, colorNode{cur.node.Right, white})
				stack = append(stack, colorNode{cur.node.Left, white})
			} else {
				result = append(result, cur.node.Val)
			}
		}
	}
	return
}

package main

import "fmt"

// 110. 平衡二叉树
// 给定一个二叉树，判断它是否是高度平衡的二叉树。
// 本题中，一棵高度平衡二叉树定义为：
//     一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
// https://leetcode-cn.com/problems/balanced-binary-tree/
func main() {
	fmt.Println(isBalanced(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	})) // FALSE
}

var depth = make(map[*TreeNode]int)

// 法一：暴力递归，从顶到底，保存中间状态
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if getAbs(getDepth(root.Left)-getDepth(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftDepth, rightDepth int
	if depth[root.Left] > 0 {
		leftDepth = depth[root.Left]
	} else {
		leftDepth = getDepth(root.Left)
		depth[root.Left] = leftDepth
	}
	if depth[root.Right] > 0 {
		rightDepth = depth[root.Right]
	} else {
		rightDepth = getDepth(root.Right)
		depth[root.Right] = rightDepth
	}
	return getMax(leftDepth, rightDepth) + 1
}

// 法二：从底到顶
func isBalanced2(root *TreeNode) bool {
	if getDepthHelper(root) < 0 {
		return false
	}
	return true
}

// 如果该节点不平衡，返回-1，其余情况返回节点深度
func getDepthHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getDepthHelper(root.Left)
	if left < 0 {
		return left
	}
	right := getDepthHelper(root.Right)
	if right < 0 {
		return right
	}
	if getAbs(left-right) <= 1 {
		return getMax(left, right) + 1
	}
	return -1
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func getAbs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

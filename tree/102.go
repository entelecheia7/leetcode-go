package main

import "fmt"

// 102. 二叉树的层序遍历
// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
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
	fmt.Println(levelOrder(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：BFS，使用一个队列，best
func levelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelNum := len(queue)
		levelVal := make([]int, levelNum)
		for i := 0; i < levelNum; i++ {
			levelVal[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[levelNum:]
		result = append(result, levelVal)
	}
	return result
}

// 法二：DFS
func levelOrder2(root *TreeNode) (result [][]int) {
	levelOrderDFSHelper(root, 0, &result)
	return result
}

// level表示节点层级，由上至下从0开始递增
func levelOrderDFSHelper(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}
	if level == len(*result) {
		*result = append(*result, []int{})
	}
	(*result)[level] = append((*result)[level], root.Val)
	levelOrderDFSHelper(root.Left, level+1, result)
	levelOrderDFSHelper(root.Right, level+1, result)
}

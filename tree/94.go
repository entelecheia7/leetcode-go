package main

import (
	"fmt"
)

// 94. 二叉树的中序遍历
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
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

	fmt.Println(inorderTraversal3(tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := inorderTraversal(root.Left)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}

// 法二：循环
func inorderTraversal2(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	cur := root
	stack := []*TreeNode{}
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}

	return result
}

// 法三：莫里斯遍历。改变树的形态，将节点以中序遍历的顺序变化成只有右子节点的二叉树
func inorderTraversal3(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	cur := root
	var pre *TreeNode
	for cur != nil {
		if cur.Left == nil {
			// 左子节点为空，处理根节点和右子树
			result = append(result, cur.Val)
			cur = cur.Right
		} else {
			// 找到左子树的最右节点pre，将根节点和右子树添加为pre的右子树
			pre = cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				pre.Right = cur
				// 重复地处理左子节点操作
				cur = cur.Left
			}
			if pre.Right == cur {
				// 还原树的结构
				result = append(result, cur.Val)
				pre.Right = nil
				cur = cur.Right
			}
		}
	}
	return result
}

type colorNode struct {
	node  *TreeNode
	color int
}

// 法四：颜色遍历，二叉树遍历通用
// 定义两种颜色，白色为首次访问，灰色为第二次，灰色节点可以直接输出
func inorderTraversal4(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}

	white, gray := 0, 1
	stack := []colorNode{{root, white}}
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.node != nil {
			if cur.color == white { // 第一次访问
				stack = append(stack, colorNode{cur.node.Right, white})
				cur.color = gray
				stack = append(stack, cur)
				stack = append(stack, colorNode{cur.node.Left, white})
			} else {
				result = append(result, cur.node.Val)
			}
		}
	}
	return result
}

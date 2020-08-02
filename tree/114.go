package main

import "github.com/davecgh/go-spew/spew"

// 114. 二叉树展开为链表
// 给定一个二叉树，原地将它展开为一个单链表。
// 例如，给定二叉树
// 1
// / \
// 2   5
// / \   \
// 3   4   6
// 将其展开为：
// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6
// https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
func main() {
	t := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	flatten(t)
	spew.Dump(t)
}

// 法一：循环
// 将左子树赋给Right，右子树添加至左子树的最右节点
func flatten(root *TreeNode) {
	for root != nil {
		if root.Left != nil {
			// 找到左子树的最右节点
			p := root.Left
			for p.Right != nil {
				p = p.Right
			}
			p.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		// 重置循环条件
		root = root.Right
	}
}

// 法二：递归
// best
func flatten2(root *TreeNode) {
	var pre *TreeNode // 在每一轮处理完成后保存根节点，作为下一轮的右子树
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Right)
		helper(root.Left)
		root.Right = pre
		root.Left = nil
		pre = root
	}
	helper(root)
}

// 法三：栈+遍历
// 每遍历一个节点，当前节点的右指针更新为上一个节点。
func flatten3(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{}
	var cur, pre *TreeNode
	cur = root
	for cur != nil || len(stack) > 0 {
		// 将根节点和右子树压栈
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}
		cur = stack[len(stack)-1]
		// 不存在左子节点 或 右子树已经处理，就处理根节点
		if cur.Left == nil || cur.Left == pre {
			stack = stack[:len(stack)-1]
			cur.Right = pre
			cur.Left = nil
			pre = cur
			cur = nil
		} else { // 优先处理左子树
			cur = cur.Left
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

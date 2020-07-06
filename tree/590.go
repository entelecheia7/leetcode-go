package main

import "fmt"

// 590. N叉树的后序遍历
// 给定一个 N 叉树，返回其节点值的后序遍历。
// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
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

	fmt.Println(postorder(tree))
	fmt.Println(postorder2(tree))
	fmt.Println(postorder3(tree))
}

type Node struct {
	Val      int
	Children []*Node
}

// 法一：递归
func postorder(root *Node) (result []int) {
	if root == nil {
		return nil
	}
	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}
	result = append(result, root.Val)
	return result
}

// 法二：颜色遍历
type colorNode struct {
	node  *Node
	color int
}

func postorder2(root *Node) (result []int) {
	if root == nil {
		return nil
	}
	white, gray := 0, 1
	stack := []colorNode{{root, white}}
	for len(stack) > 0 {
		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if back.color == white {
			back.color = gray
			stack = append(stack, back)
			for i := len(back.node.Children) - 1; i >= 0; i-- {
				stack = append(stack, colorNode{back.node.Children[i], white})
			}
		} else {
			result = append(result, back.node.Val)
		}
	}
	return
}

// 法三：前序遍历，再将结果倒序
func postorder3(root *Node) (result []int) {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		back := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, back.Val)
		for _, c := range back.Children {
			stack = append(stack, c)
		}
	}
	// reverse
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return
}

package main

import "fmt"

// 337. 打家劫舍 III
// 在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
// 计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。
// https://leetcode-cn.com/problems/house-robber-iii/
func main() {
	fmt.Println(rob2(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 1,
			},
		},
	})) // 7
}

// 法一：暴力递归
// 当前层要么选择，要么不选择
func rob(root *TreeNode) int {
	// 处于当前节点所能取得的最大和
	m := make(map[*TreeNode]int)
	return robHelper(root, m)
}

// available代表当前层是否可选
func robHelper(root *TreeNode, m map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if m[root] > 0 {
		return m[root]
	}
	// r1代表选择当前层，r2代表跳过当前层
	r1 := root.Val
	if root.Left != nil {
		r1 += robHelper(root.Left.Left, m) + robHelper(root.Left.Right, m)
	}
	if root.Right != nil {
		r1 += robHelper(root.Right.Left, m) + robHelper(root.Right.Right, m)
	}

	r2 := robHelper(root.Left, m) + robHelper(root.Right, m)

	r := getMax(r1, r2)
	m[root] = r
	return r
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 法二：另一种写法
// best
func rob2(root *TreeNode) int {
	r := robHelper2(root)
	return getMax(r[0], r[1])
}

// 返回 含root的最大子树和 和 不含root的最大子树和
func robHelper2(root *TreeNode) (result [2]int) {
	if root == nil {
		return
	}
	left := robHelper2(root.Left)
	right := robHelper2(root.Right)

	result[0] = root.Val + left[1] + right[1]
	result[1] = getMax(left[0], left[1]) + getMax(right[0], right[1])

	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 133. 克隆图
// 给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
// 图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。
// class Node {
//     public int val;
//     public List<Node> neighbors;
// }
// 测试用例格式：
// 简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。该图在测试用例中使用邻接列表表示。
// 邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。
// 给定节点将始终是图中的第一个节点（值为 1）。你必须将 给定节点的拷贝 作为对克隆图的引用返回。
// 提示：
// 节点数不超过 100 。
// 每个节点值 Node.val 都是唯一的，1 <= Node.val <= 100。
// 无向图是一个简单图，这意味着图中没有重复的边，也没有自环。
// 由于图是无向的，如果节点 p 是节点 q 的邻居，那么节点 q 也必须是节点 p 的邻居。
// 图是连通图，你可以从给定节点访问到所有节点。
// https://leetcode-cn.com/problems/clone-graph
func main() {
	v1 := newNode(1)
	v2 := newNode(2)
	v1.Neighbors = append(v1.Neighbors, v2)
	v2.Neighbors = append(v2.Neighbors, v1)
	spew.Dump(cloneGraph(v1))
}

// 法一：BFS
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := map[*Node]*Node{} // origNode => copierNode
	queue := []*Node{node}       // origNode
	visited[node] = newNode(node.Val)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, n := range cur.Neighbors {
			if visited[n] == nil {
				visited[n] = newNode(n.Val)
				queue = append(queue, n)
			}
			visited[cur].Neighbors = append(visited[cur].Neighbors, visited[n])
		}
	}
	return visited[node]
}
func newNode(val int) *Node {
	return &Node{Val: val}
}

// 法二：DFS
func cloneGraph2(node *Node) *Node {
	visited := make(map[*Node]*Node, 101)
	return dfsHelper(node, visited)
}
func dfsHelper(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, exist := visited[node]; exist {
		return v
	}
	copier := newNode(node.Val)
	visited[node] = copier
	for _, neighbor := range node.Neighbors {
		copier.Neighbors = append(copier.Neighbors, dfsHelper(neighbor, visited))
	}
	return copier
}

type Node struct {
	Val       int
	Neighbors []*Node
}

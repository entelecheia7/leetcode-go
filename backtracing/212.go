package main

import (
	"fmt"
	// "github.com/davecgh/go-spew/spew"
)

// 212. 单词搜索 II
// 给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
// 说明:
// 你可以假设所有输入都由小写字母 a-z 组成。
// 提示:
//     你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
//     如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？ 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。
// https://leetcode-cn.com/problems/word-search-ii
func main() {
	// fmt.Println(findWords([][]byte{
	// 	{'o', 'a', 'a', 'n'},
	// 	{'e', 't', 'a', 'e'},
	// 	{'i', 'h', 'k', 'r'},
	// 	{'i', 'f', 'l', 'v'},
	// }, []string{"oath", "pea", "eat", "rain"})) // [eat oath]

	fmt.Println(findWords([][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
		{'a', 'a', 'a', 'a'},
	}, []string{"aaaaaaaaaaaa", "aaaaaaaaaaaaa", "aaaaaaaaaaab"})) // ["aaaaaaaaaaaa"]

}

// TRIE树适合查找公共前缀匹配的字符串
// 将 words构建成一棵TRIE树
// 然后使用回溯进行查找
func findWords(board [][]byte, words []string) (result []string) {
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	trie := buildTrieTree(words)
	find := make(map[string]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if trie.next[board[i][j]-'a'] != nil {
				backTracing(board, []byte{}, i, j, m, n, trie, find)
			}
		}
	}
	for k := range find {
		result = append(result, k)
	}
	return
}

var diff [4][2]int = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// 为了去重，结果集使用一个map进行保存
func backTracing(board [][]byte, cur []byte, i, j, m, n int, trie *TrieNode, find map[string]bool) {
	if i < 0 || i == m || j < 0 || j == n || board[i][j] == '.' {
		return
	}
	k := board[i][j] - 'a'
	if trie.next[k] == nil {
		return
	}
	cur = append(cur, board[i][j])
	tmp := board[i][j]
	board[i][j] = '.'
	if trie.next[k].isEnd {
		find[string(cur)] = true
	}
	for _, v := range diff {
		backTracing(board, cur, i+v[0], j+v[1], m, n, trie.next[k], find)
	}
	cur = cur[:len(cur)-1]
	board[i][j] = tmp
}

func buildTrieTree(words []string) *TrieNode {
	trie := &TrieNode{}
	for _, word := range words {
		p := trie
		for _, char := range word {
			c := char - 'a'
			if p.next[c] == nil {
				p.next[c] = &TrieNode{}
			}
			p = p.next[c]
		}
		p.isEnd = true
	}
	return trie
}

type TrieNode struct {
	isEnd bool
	next  [26]*TrieNode
}

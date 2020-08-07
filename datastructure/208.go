package main

import (
	"fmt"
)

// 208. 实现 Trie (前缀树)
// 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
// 说明:
//     你可以假设所有的输入都是由小写字母 a-z 构成的。
//     保证所有输入均为非空字符串。
// https://leetcode-cn.com/problems/implement-trie-prefix-tree/
func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("app"))     // false
	fmt.Println(trie.StartsWith("app")) // true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // true
}

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, char := range word {
		c := char - 'a'
		if this.next[c] == nil {
			this.next[c] = &Trie{}
		}
		this = this.next[c]
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, char := range word {
		c := char - 'a'
		if this.next[c] == nil {
			return false
		}
		this = this.next[c]
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, char := range prefix {
		c := char - 'a'
		if this.next[c] == nil {
			return false
		}
		this = this.next[c]
	}
	return true
}

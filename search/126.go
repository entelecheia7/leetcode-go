package main

import (
	"fmt"
)

// 126. 单词接龙 II
// 给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：
// 每次转换只能改变一个字母。
// 转换后得到的单词必须是字典中的单词。
// 说明:
// 如果不存在这样的转换序列，返回一个空列表。
// 所有单词具有相同的长度。
// 所有单词只由小写字母组成。
// 字典中不存在重复的单词。
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
// https://leetcode-cn.com/problems/word-ladder-ii/description/
func main() {
	// [
	// 	["hit","hot","dot","dog","cog"],
	// 	["hit","hot","lot","log","cog"]
	// ]
	// fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))

	// fmt.Println(findLadders("a", "c", []string{"a", "b", "c"}))

	// [["red","ted","tad","tax"],["red","ted","tex","tax"],["red","rex","tex","tax"]]
	// fmt.Println(findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))

	fmt.Println(findLadders("qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}))
}

// 法一：BFS
func findLadders(beginWord string, endWord string, wordList []string) (result [][]string) {
	wordN := len(wordList)
	if wordN == 0 {
		return nil
	}
	n := len(beginWord)
	wordListMap := make(map[string]int, wordN)
	for key, word := range wordList {
		wordListMap[word] = key
	}
	if wordListMap[endWord] == 0 && wordList[0] != endWord { // endword不在字典中
		return nil
	}
	// 将word的转换关系抽象为无向图
	if wordListMap[beginWord] == 0 && wordList[0] != beginWord { // 保证beginword也在字典中
		wordList = append(wordList, beginWord)
		wordListMap[beginWord] = wordN
		wordN++
	}
	graph := make([][]int, len(wordList))
	for i := 0; i < wordN-1; i++ {
		for j := i + 1; j < wordN; j++ {
			if checkConvertion(wordList[i], wordList[j], n) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}
	// fmt.Println("graph:", graph, wordListMap)

	path := map[int][][]int{wordListMap[beginWord]: [][]int{{wordListMap[beginWord]}}} // 保存从start到key的可能路径
	queue := []int{wordListMap[beginWord]}
	visited := make(map[int]bool)
	addToQueue := map[int]bool{wordListMap[beginWord]: true}
	level := 0
Loop:
	for len(queue) > 0 {
		level++
		size := len(queue)
		fmt.Println("queue:", len(queue))
		for i := 0; i < size; i++ {
			cur := queue[i]
			if visited[cur] {
				continue
			}
			if cur == wordListMap[endWord] {
				break Loop
			}
			visited[cur] = true
			next := graph[queue[i]]
			for _, p := range path[cur] { // previous path
				for _, n := range next { // next path
					if visited[n] {
						continue
					}
					tmp := make([]int, len(p)+1)
					copy(tmp, p)
					tmp[len(p)] = n
					path[n] = append(path[n], tmp)
					if !addToQueue[n] {
						queue = append(queue, n)
						addToQueue[n] = true
					}
				}
			}
		}
		printPath(wordList, path)
		queue = queue[size:]
	}

	if len(path[wordListMap[endWord]]) == 0 {
		return nil
	}
	for _, p := range path[wordListMap[endWord]] {
		if len(p) == level {
			tmp := make([]string, len(p))
			for i := 0; i < len(p); i++ {
				tmp[i] = wordList[p[i]]
			}
			result = append(result, tmp)
		}
	}

	return result
}

// 检查从start是否可以更改1个字母转换为end
// n是字符串长度
func checkConvertion(start, end string, n int) bool {
	var k byte
	word := []byte(start)
	for i := 0; i < n; i++ {
		old := word[i]
		for k = 'a'; k <= 'z'; k++ {
			if k != old {
				word[i] = k
				if string(word) == end {
					return true
				}
			}
		}
		word[i] = old
	}
	return false
}

func printPath(wordList []string, path map[int][][]int) {
	s := ""
	for k, p := range path {
		s += fmt.Sprintf("%s: ", wordList[k])
		for _, v := range p {
			for i := 0; i < len(v); i++ {
				s += fmt.Sprintf("%s-", wordList[v[i]])
			}
			s += fmt.Sprintf("、")
		}
	}
	fmt.Println(s + "\n=====\n")
}

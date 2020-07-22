package main

import (
	"fmt"
	"math"
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
	// fmt.Println(findLadders2("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))

	// fmt.Println(findLadders2("a", "c", []string{"a", "b", "c"}))

	// [["red","ted","tad","tax"],["red","ted","tex","tax"],["red","rex","tex","tax"]]
	fmt.Println(findLadders2("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))

	// fmt.Println(findLadders2("qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}))
}

// 法一：单向BFS
// 构建了一个图，在图的基础上进行BFS，queue中保存转换路径
// 代码复杂度高，时间和空间复杂度较差
func findLadders(beginWord string, endWord string, wordList []string) (result [][]string) {
	wordN := len(wordList)
	if wordN == 0 {
		return nil
	}
	n := len(beginWord)
	wordListMap := make(map[string]int, wordN) // key是string，value是字符串在wordList的下标
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

	queue := [][]string{{beginWord}} // queue保存的是从startword扩散的路径
	cost := make([]int, wordN)       // 记录从startword转为wordList[i]的最短长度
	for i := 0; i < wordN; i++ {
		cost[i] = math.MaxInt64
	}
	cost[wordListMap[beginWord]] = 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			last := cur[len(cur)-1]
			if last == endWord {
				// add to result
				tmp := make([]string, len(cur))
				copy(tmp, cur)
				result = append(result, tmp)
			} else {
				for _, next := range graph[wordListMap[last]] {
					if cost[next] >= cost[wordListMap[last]]+1 { // 保证转换路径不回退
						cost[next] = cost[wordListMap[last]] + 1
						tmp := make([]string, len(cur)+1)
						copy(tmp, cur)
						tmp[len(cur)] = wordList[next]
						queue = append(queue, tmp)
					}
				}
			}
		}
		if len(result) > 0 {
			break
		}

		queue = queue[size:]
	}

	return result
}

// 检查从start是否可以更改1个字母转换为end
// n是字符串长度
func checkConvertion(start, end string, n int) bool {
	for i := 0; i < n; i++ {
		if start[i] != end[i] {
			return start[i+1:] == end[i+1:]
		}
	}
	return false // start == end
}

// 法二：双向BFS
// best
func findLadders2(beginWord string, endWord string, wordList []string) (result [][]string) {
	if len(wordList) == 0 {
		return nil
	}
	wordListMap := make(map[string]bool, len(wordList)) // 一个全局的未访问元素数组
	for _, w := range wordList {
		wordListMap[w] = true
	}
	if !wordListMap[endWord] {
		return nil
	}
	delete(wordListMap, endWord)

	// queue保存这一层的节点
	queue, queueFromEnd := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	n := len(beginWord)
	var j byte
	endFlag, reverseFlag := false, false
	path := make(map[string][]string) // 记录key可以转换的单次

	for len(queue) > 0 && len(queueFromEnd) > 0 && !endFlag {
		if len(queue) > len(queueFromEnd) {
			queue, queueFromEnd = queueFromEnd, queue
			reverseFlag = !reverseFlag
		}
		for w := range queue {
			delete(wordListMap, w)
		}
		newqueue := make(map[string]bool)
		for word := range queue {
			tmp := []byte(word)
			for i := 0; i < n; i++ {
				old := tmp[i]
				for j = 'a'; j <= 'z'; j++ {
					if j != old {
						tmp[i] = j
						convertion := string(tmp)
						if queueFromEnd[convertion] { // 双向BFS相遇
							if reverseFlag {
								path[convertion] = append(path[convertion], word)
							} else {
								path[word] = append(path[word], convertion)
							}
							endFlag = true
						} else if wordListMap[convertion] { // 未访问过，说明到达下一层
							newqueue[convertion] = true
							if reverseFlag {
								path[convertion] = append(path[convertion], word)
							} else {
								path[word] = append(path[word], convertion)
							}
						}
					}
				}
				tmp[i] = old
			}
		}
		queue = newqueue
	}

	// DFS，从beginWord开始组装结果
	cur := []string{beginWord}
	var generator func([]string)
	generator = func(words []string) {
		for _, n := range words {
			cur = append(cur, n)
			if n == endWord {
				tmp := make([]string, len(cur))
				copy(tmp, cur)
				result = append(result, tmp)
			} else {
				generator(path[n])
			}
			cur = cur[:len(cur)-1]
		}

	}
	generator(path[beginWord])

	return result
}

/*
LeetCode 720: https://leetcode.com/problems/longest-word-in-dictionary/description/
*/

package leetcode

import "fmt"

func longestWord(words []string) string {
	root := buildTrie(words)
	return findLongestWord(root, "")
}

func buildTrie(words []string) *trieNode {
	root := &trieNode{
		children: [26]*trieNode{},
	}

	for _, word := range words {
		node := root
		for _, ch := range word {
			index := int(ch - 'a')
			if node.children[index] == nil {
				node.children[index] = &trieNode{
					children: [26]*trieNode{},
				}
			}

			node = node.children[index]
		}

		node.isWord = true
	}

	return root
}

func findLongestWord(root *trieNode, word string) string {
	longest := word
	for i, child := range root.children {
		next := word
		if child != nil && child.isWord {
			next = findLongestWord(child, word+fmt.Sprintf("%c", i+'a'))
		}

		if len(next) > len(longest) {
			longest = next
		}
	}

	return longest
}

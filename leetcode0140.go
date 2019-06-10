/*
LeetCode 140: https://leetcode.com/problems/word-break-ii/
*/

package leetcode

import (
	"fmt"
)

type trieNode struct {
	children [26]*trieNode
	isWord   bool
}

func wordBreak(s string, wordDict []string) []string {
	root := &trieNode{}
	for _, word := range wordDict {
		node := root
		for _, ch := range word {
			index := int(ch - 'a')
			if node.children[index] == nil {
				node.children[index] = &trieNode{}
			}

			node = node.children[index]
		}

		node.isWord = true
	}

	memo := make([][]string, len(s)+1)
	return helper140(s, 0, root, root, "", memo)
}

func helper140(s string, index int, root, node *trieNode, word string, memo [][]string) []string {
	result := make([]string, 0)
	if index == len(s) && node == root {
		result = append(result, "")
	} else if index < len(s) && node != nil {
		ch := s[index]
		word, node = fmt.Sprintf("%s%c", word, ch), node.children[int(ch-'a')]
		if node != nil && node.isWord {
			var next []string
			if memo[index+1] != nil {
				next = memo[index+1]
			} else {
				next = helper140(s, index+1, root, root, "", memo)
				memo[index+1] = next
			}

			for _, n := range next {
				if n == "" {
					result = append(result, word)
				} else {
					result = append(result, fmt.Sprintf("%s %s", word, n))
				}
			}
		}

		result = append(result, helper140(s, index+1, root, node, word, memo)...)
	}

	return result
}

/*
LeetCode 212: https://leetcode.com/problems/word-search-ii/
*/

package leetcode

import (
	"fmt"
)

// NOTE: trieNode is defined in commontyeps.go.
func findWords(board [][]byte, words []string) []string {
	root := &trieNode{}
	for _, word := range words {
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

	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	set := make(map[string]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			helper212(board, i, j, visited, root, "", set)
		}
	}

	result := make([]string, 0)
	for s := range set {
		result = append(result, s)
	}

	return result
}

func helper212(board [][]byte, i, j int, visited [][]bool, node *trieNode, word string, set map[string]bool) {
	word = fmt.Sprintf("%s%c", word, board[i][j])
	visited[i][j] = true
	node = node.children[int(board[i][j]-'a')]
	if node != nil && node.isWord {
		set[word] = true
	}

	if node != nil {
		dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, dir := range dirs {
			r, c := i+dir[0], j+dir[1]
			if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) && !visited[r][c] {
				helper212(board, r, c, visited, node, word, set)
			}
		}
	}

	visited[i][j] = false
}

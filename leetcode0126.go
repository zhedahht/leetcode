/*
LeetCode 126: https://leetcode.com/problems/word-ladder-ii/
*/

package leetcode

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	visited := make(map[string]bool)
	for _, word := range wordList {
		visited[word] = false
	}

	prevWordsMap := make(map[string][]string)
	level1, level2 := make(map[string]bool), make(map[string]bool)
	level1[beginWord] = true

outer:
	for len(level1) > 0 {
		for word := range level1 {
			if word == endWord {
				break outer
			}

			nextWords := getNextWords126(word, visited)
			for _, next := range nextWords {
				if _, exists := prevWordsMap[next]; !exists {
					prevWordsMap[next] = make([]string, 0)
				}

				prevWordsMap[next] = append(prevWordsMap[next], word)
				level2[next] = true
			}
		}

		for next := range level2 {
			visited[next] = true
		}

		level1, level2 = level2, make(map[string]bool)
	}

	result := make([][]string, 0)
	if _, exists := prevWordsMap[endWord]; exists {
		result = helper126(endWord, beginWord, prevWordsMap, make([]string, 0), result)
	}

	return result
}

func getNextWords126(word string, visited map[string]bool) []string {
	result := make([]string, 0)
	bytes := []byte(word)
	for i, b := range bytes {
		for j := byte('z'); j >= byte('a'); j-- {
			if j != b {
				bytes[i] = j
				candidate := string(bytes)
				if v, exists := visited[candidate]; exists && !v {
					result = append(result, candidate)
				}
			}
		}

		bytes[i] = b
	}

	return result
}

func helper126(word, beginWord string, prevWordsMap map[string][]string, path []string, result [][]string) [][]string {
	path = append(path, word)
	if word == beginWord {
		clone := make([]string, len(path))
		copy(clone, path)
		for i, j := 0, len(clone)-1; i < j; i, j = i+1, j-1 {
			clone[i], clone[j] = clone[j], clone[i]
		}

		result = append(result, clone)
	} else {
		for _, prev := range prevWordsMap[word] {
			result = helper126(prev, beginWord, prevWordsMap, path, result)
		}
	}

	path = path[:len(path)-1]
	return result
}

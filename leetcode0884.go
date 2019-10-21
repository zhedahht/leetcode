/*
LeetCode 884: https://leetcode.com/problems/uncommon-words-from-two-sentences/
*/

package leetcode

import "strings"

func uncommonFromSentences(A string, B string) []string {
	buildMap := func(words []string) map[string]int {
		strToCount := make(map[string]int)
		for _, word := range words {
			count := strToCount[word]
			strToCount[word] = count + 1
		}

		return strToCount
	}

	words1, words2 := strings.Split(A, " "), strings.Split(B, " ")
	map1, map2 := buildMap(words1), buildMap(words2)
	var result []string
	for _, word := range words1 {
		if map1[word] == 1 && map2[word] == 0 {
			result = append(result, word)
		}
	}

	for _, word := range words2 {
		if map2[word] == 1 && map1[word] == 0 {
			result = append(result, word)
		}
	}

	return result
}

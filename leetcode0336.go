/*
LeetCode 336: https://leetcode.com/problems/palindrome-pairs/
*/

package leetcode

import "fmt"

func palindromePairs(words []string) [][]int {
	isPalin := func(word string) bool {
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			if word[i] != word[j] {
				return false
			}
		}

		return true
	}

	reverse := func(word string) string {
		bytes := []byte(word)
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}

		return string(bytes)
	}

	original := make(map[string]int)
	for i, word := range words {
		original[word] = i
	}

	result := make([][]int, 0)
	found := make(map[string]bool)
	for index, word := range words {
		for i := 0; i <= len(word); i++ {
			prefix := word[:i]
			suffix := word[i:]

			rev := reverse(suffix)
			revIndex, exists := original[rev]
			key := fmt.Sprintf("%d,%d", original[rev], index)
			if isPalin(prefix) && exists && revIndex != index && !found[key] {
				result = append(result, []int{original[rev], index})
				found[key] = true
			}

			rev = reverse(prefix)
			revIndex, exists = original[rev]
			key = fmt.Sprintf("%d,%d", index, original[rev])
			if isPalin(suffix) && exists && revIndex != index && !found[key] {
				result = append(result, []int{index, original[rev]})
				found[key] = true
			}
		}
	}

	return result
}

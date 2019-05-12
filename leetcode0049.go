/*
LeetCode 49: https://leetcode.com/problems/group-anagrams/
*/

package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	sortString := func(str string) string {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})

		return string(bytes)
	}

	sortedToAnagrams := make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		anagrams, exists := sortedToAnagrams[sorted]
		if !exists {
			anagrams = make([]string, 0)
		}

		sortedToAnagrams[sorted] = append(anagrams, str)
	}

	result := make([][]string, 0)
	for _, anagrams := range sortedToAnagrams {
		result = append(result, anagrams)
	}

	return result
}

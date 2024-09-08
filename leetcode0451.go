/*
LeetCode 451: https://leetcode.com/problems/sort-characters-by-frequency/description/
*/

package leetcode

import (
	"sort"
	"strings"
)

type CharCount struct {
	char  rune
	count int
}

func frequencySort(s string) string {
	counts := make(map[rune]int)
	for _, ch := range s {
		counts[ch]++
	}

	charCounts := make([]CharCount, 0)
	for k, v := range counts {
		charCounts = append(charCounts, CharCount{
			char:  k,
			count: v,
		})
	}

	sort.Slice(charCounts, func(i, j int) bool {
		return charCounts[i].count > charCounts[j].count
	})

	var sb strings.Builder
	for _, charCount := range charCounts {
		for i := 0; i < charCount.count; i++ {
			sb.WriteRune(charCount.char)
		}
	}

	return sb.String()
}

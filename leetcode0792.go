/*
LeetCode 792: https://leetcode.com/problems/number-of-matching-subsequences/
*/

package leetcode

func numMatchingSubseq(s string, words []string) int {
	chToPos := make(map[rune][]int)
	for i, ch := range s {
		chToPos[ch] = append(chToPos[ch], i)
	}

	result := 0
	for _, word := range words {
		if isMatching(chToPos, word) {
			result++
		}
	}
	return result
}

func isMatching(chToPos map[rune][]int, word string) bool {
	i, j := 0, 0
	for _, ch := range word {
		next := findNext(chToPos, ch, i)
		if next < 0 {
			break
		}

		i, j = next+1, j+1
	}

	return j == len(word)
}

func findNext(chToPos map[rune][]int, ch rune, val int) int {
	vals, exists := chToPos[ch]
	if !exists {
		return -1
	}

	left, right := 0, len(vals)-1
	for left <= right {
		mid := (left + right) / 2
		if vals[mid] < val {
			left = mid + 1
		} else {
			if mid == 0 || vals[mid-1] < val {
				return vals[mid]
			}
			right = mid - 1
		}
	}

	return -1
}

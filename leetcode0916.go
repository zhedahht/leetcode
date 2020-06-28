/*
LeetCode 916: https://leetcode.com/problems/word-subsets/
*/

package leetcode

func wordSubsets(A []string, B []string) []string {
	getCharCounts := func(str string) map[rune]int {
		counts := make(map[rune]int)
		for _, r := range str {
			counts[r] = counts[r] + 1
		}

		return counts
	}

	maxCounts := make(map[rune]int, 0)
	for _, str := range B {
		countsB := getCharCounts(str)
		for r, c := range countsB {
			if maxCounts[r] < c {
				maxCounts[r] = c
			}
		}
	}

	result := make([]string, 0)
	for _, str := range A {
		countsA := getCharCounts(str)
		isValid := true

		for r, c := range maxCounts {
			if countsA[r] < c {
				isValid = false
				break
			}
		}

		if isValid {
			result = append(result, str)
		}
	}

	return result
}

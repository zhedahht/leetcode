/*
LeetCode 243: https://leetcode.com/problems/shortest-word-distance/
*/

package leetcode

import "math"

func shortestDistance(words []string, word1 string, word2 string) int {
	indices1, indices2 := make([]int, 0), make([]int, 0)
	for i, word := range words {
		if word == word1 {
			indices1 = append(indices1, i)
		} else if word == word2 {
			indices2 = append(indices2, i)
		}
	}

	dist := len(words)
	i, j := 0, 0
	for i < len(indices1)-1 || j < len(indices2)-1 {
		dist = int(math.Min(float64(dist), math.Abs(float64(indices1[i]-indices2[j]))))
		if j == len(indices2)-1 || (i < len(indices1)-1 && indices1[i] < indices2[j]) {
			i++
		} else {
			j++
		}
	}

	dist = int(math.Min(float64(dist), math.Abs(float64(indices1[i]-indices2[j]))))
	return dist
}

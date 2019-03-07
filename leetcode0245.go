/*
LeetCode 245: https://leetcode.com/problems/shortest-word-distance-iii/
*/

package leetcode

import "math"

func shortestWordDistance(words []string, word1 string, word2 string) int {
	dist := len(words)
	if word1 != word2 {
		indices1, indices2 := make([]int, 0), make([]int, 0)
		for i, word := range words {
			if word == word1 {
				indices1 = append(indices1, i)
			} else if word == word2 {
				indices2 = append(indices2, i)
			}
		}

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
	} else {
		indices1 := make([]int, 0)
		for i, word := range words {
			if word == word1 {
				indices1 = append(indices1, i)
			}
		}

		for i := 1; i < len(indices1); i++ {
			dist = int(math.Min(float64(dist), math.Abs(float64(indices1[i]-indices1[i-1]))))
		}
	}

	return dist
}

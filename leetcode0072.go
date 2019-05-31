/*
LeetCode 72: https://leetcode.com/problems/edit-distance/
*/

package leetcode

import (
	"math"
)

func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = i
	}

	for i := 0; i < len(word1); i++ {
		prev := i
		dp[0] = i + 1
		for j := 0; j < len(word2); j++ {
			cur := prev
			if word1[i] != word2[j] {
				min := math.Min(float64(dp[j+1]), float64(dp[j]))
				cur = int(math.Min(min, float64(prev))) + 1
			}

			prev, dp[j+1] = dp[j+1], cur
		}
	}

	return dp[len(word2)]
}

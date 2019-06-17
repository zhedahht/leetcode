/*
LeetCode 403: https://leetcode.com/problems/frog-jump/
*/

package leetcode

func canCross(stones []int) bool {
	dp := make([]map[int]bool, len(stones))
	for i := range stones {
		dp[i] = make(map[int]bool)
	}

	if stones[1] == 1 {
		dp[1][1] = true
	}

	for i := 2; i < len(stones); i++ {
		for j := i - 1; j >= 0; j-- {
			if len(dp[j]) > 0 {
				dist := stones[i] - stones[j]
				dists := [3]int{dist, dist - 1, dist + 1}
				for _, d := range dists {
					if _, exists := dp[j][d]; exists {
						dp[i][dist] = true
						break
					}
				}
			}
		}
	}

	return len(dp[len(stones)-1]) > 0
}

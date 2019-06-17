/*
LeetCode 403: https://leetcode.com/problems/frog-jump/
*/

package leetcode

func canCross(stones []int) bool {
	dp, stoneToIndex := make([][]bool, len(stones)), make(map[int]int)
	for i, s := range stones {
		dp[i] = make([]bool, len(stones))
		stoneToIndex[s] = i
	}

	if stones[1] == 1 {
		dp[0][1] = true
	}

	for i := 1; i < len(stones)-1; i++ {
		for j := 0; j < i; j++ {
			if dp[j][i] {
				dist := stones[i] - stones[j]
				dists := [3]int{dist, dist - 1, dist + 1}
				for _, d := range dists {
					if d > 0 {
						pos := stones[i] + d
						if next, exists := stoneToIndex[pos]; exists {
							dp[i][next] = true
						}
					}
				}
			}
		}
	}

	for i := 0; i < len(stones); i++ {
		if dp[i][len(stones)-1] {
			return true
		}
	}

	return false
}

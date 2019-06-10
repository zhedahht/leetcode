/*
LeetCode 174: https://leetcode.com/problems/dungeon-game/
*/

package leetcode

import (
	"math"
)

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}

	rows, cols := len(dungeon), len(dungeon[0])
	dp := make([]int, cols)
	left := 1
	for i := cols - 1; i >= 0; i-- {
		dp[i] = int(math.Max(1.0, float64(left-dungeon[rows-1][i])))
		left = dp[i]
	}

	for i := rows - 2; i >= 0; i-- {
		left := 0x7fffffff
		for j := cols - 1; j >= 0; j-- {
			valRight := int(math.Max(1.0, float64(left-dungeon[i][j])))
			valDown := int(math.Max(1.0, float64(dp[j]-dungeon[i][j])))
			dp[j] = int(math.Min(float64(valRight), float64(valDown)))
			left = dp[j]
		}
	}

	return dp[0]
}

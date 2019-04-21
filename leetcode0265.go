/*
LeetCode 265: https://leetcode.com/problems/paint-house-ii/
*/

package leetcode

import "math"

func minCostII(costs [][]int) int {
	if len(costs) == 0 || len(costs[0]) == 0 {
		return 0
	}

	minCosts := make([][]int, 2)
	minCosts[0] = make([]int, len(costs[0]))
	minCosts[1] = make([]int, len(costs[0]))

	for j := 0; j < len(costs[0]); j++ {
		minCosts[0][j] = costs[0][j]
	}

	for i := 1; i < len(costs); i++ {
		for j := 0; j < len(costs[0]); j++ {
			minCost := 0x7fffffff
			for k := 0; k < len(costs[0]); k++ {
				if k != j && minCost > minCosts[(i-1)%2][k]+costs[i][j] {
					minCost = minCosts[(i-1)%2][k] + costs[i][j]
				}
			}

			minCosts[i%2][j] = minCost
		}
	}

	result := minCosts[(len(costs)-1)%2][0]
	for i := 1; i < len(minCosts[0]); i++ {
		result = int(math.Min(float64(result), float64(minCosts[(len(costs)-1)%2][i])))
	}

	return result
}

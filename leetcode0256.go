/*
LeetCode 256: https://leetcode.com/problems/paint-house/
*/

package leetcode

import "math"

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	cost0, cost1, cost2 := costs[0][0], costs[0][1], costs[0][2]
	for i := 1; i < len(costs); i++ {
		newCost0 := costs[i][0] + int(math.Min(float64(cost1), float64(cost2)))
		newCost1 := costs[i][1] + int(math.Min(float64(cost0), float64(cost2)))
		newCost2 := costs[i][2] + int(math.Min(float64(cost0), float64(cost1)))

		cost0, cost1, cost2 = newCost0, newCost1, newCost2
	}

	return int(math.Min(float64(cost0), math.Min(float64(cost1), float64(cost2))))
}

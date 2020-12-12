/*
LeetCode 134: https://leetcode.com/problems/gas-station/
*/

package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	delta, last := 0, 0
	totalCost, totalGas := 0, 0
	for i := range gas {
		totalCost += cost[i]
		totalGas += gas[i]
		delta += gas[i] - cost[i]
		if delta < 0 {
			last = i + 1
			delta = 0
		}
	}

	if totalCost > totalGas {
		return -1
	}

	return last
}

/*
LeetCode 475: https://leetcode.com/problems/heaters/description/
*/

package leetcode

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Slice(houses, func(i, j int) bool {
		return houses[i] < houses[j]
	})
	sort.Slice(heaters, func(i, j int) bool {
		return heaters[i] < heaters[j]
	})

	newHeaters := []int{math.MinInt32}
	newHeaters = append(newHeaters, heaters...)
	newHeaters = append(newHeaters, math.MaxInt32)

	index := 1
	minRadius := 0
	for _, house := range houses {
		for newHeaters[index] < house {
			index++
		}

		radius := min(newHeaters[index]-house, house-newHeaters[index-1])
		minRadius = max(minRadius, radius)
	}

	return minRadius
}

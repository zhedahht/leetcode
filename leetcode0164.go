/*
LeetCode 164: https://leetcode.com/problems/maximum-gap/
*/

package leetcode

import "math"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	min, max := nums[0], nums[0]
	for _, num := range nums {
		min = int(math.Min(float64(min), float64(num)))
		max = int(math.Max(float64(max), float64(num)))
	}

	length := float64(max-min+1) / float64(len(nums))
	mins, maxs, hits := make([]int, len(nums)), make([]int, len(nums)), make([]bool, len(nums))
	for i := range nums {
		mins[i], maxs[i] = 0x7fffffff, 0
	}

	for _, num := range nums {
		index := int(float64(num-min) / length)
		mins[index] = int(math.Min(float64(mins[index]), float64(num)))
		maxs[index] = int(math.Max(float64(maxs[index]), float64(num)))
		hits[index] = true
	}

	result, i := 0, 0
	for ; i < len(nums) && !hits[i]; i++ {
	}

	min, max = mins[i], maxs[i]
	for i++; i < len(nums); i++ {
		if hits[i] {
			result = int(math.Max(float64(result), float64(mins[i]-max)))
			min, max = mins[i], maxs[i]
		}
	}

	return result
}

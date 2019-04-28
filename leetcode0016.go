/*
LeetCode 16: https://leetcode.com/problems/3sum-closest/
*/

package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := (nums[0] + nums[1] + nums[2]) - target
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return target
			}

			if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
				closest = sum
			}

			if sum > target {
				k--
			} else {
				j++
			}
		}
	}

	return closest
}

/*
LeetCode 628: https://leetcode.com/problems/maximum-product-of-three-numbers/description/
*/

package leetcode

import (
	"math"
	"slices"
)

func maximumProduct(nums []int) int {
	slices.Sort(nums)

	max := math.MinInt
	for i := 0; i < len(nums)-1; i++ {
		prod := nums[i] * nums[i+1]
		if prod >= 0 {
			if i < len(nums)-2 {
				prod *= nums[len(nums)-1]
			} else {
				prod *= nums[i-1]
			}
		} else {
			if i == 0 {
				prod *= nums[2]
			} else {
				prod *= nums[0]
			}
		}

		if prod > max {
			max = prod
		}
	}

	return max
}

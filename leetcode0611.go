/*
LeetCode 611: https://leetcode.com/problems/valid-triangle-number/description/
*/

package leetcode

import "slices"

func triangleNumber(nums []int) int {
	slices.Sort(nums)
	count := 0
	k := 0
	for i := 0; i < len(nums); i++ {
		k = i + 2
		for j := i + 1; j < len(nums); j++ {
			if j+1 > k {
				k = j + 1
			}
			for k < len(nums) && nums[k] < nums[i]+nums[j] {
				k++
			}

			count += (k - j - 1)
		}
	}

	return count
}

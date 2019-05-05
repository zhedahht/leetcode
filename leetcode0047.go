/*
LeetCode 47: https://leetcode.com/problems/permutations-ii/
*/

package leetcode

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return helper47(nums, 0, make([][]int, 0))
}

func helper47(nums []int, i int, result [][]int) [][]int {
	if i == len(nums) {
		clone := make([]int, len(nums))
		copy(clone, nums)
		result = append(result, clone)
	} else {
		swapped := make(map[int]bool)
		for j := i; j < len(nums); j++ {
			if _, ok := swapped[nums[j]]; !ok {
				swapped[nums[j]] = true

				nums[i], nums[j] = nums[j], nums[i]
				result = helper47(nums, i+1, result)
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return result
}

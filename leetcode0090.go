/*
LeetCode 90: https://leetcode.com/problems/subsets-ii/
*/

package leetcode

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return helper90(nums, 0, make([]int, 0), make([][]int, 0))
}

func helper90(nums []int, i int, comb []int, result [][]int) [][]int {
	if i == len(nums) {
		clone := make([]int, len(comb))
		copy(clone, comb)
		result = append(result, clone)
	} else {
		comb = append(comb, nums[i])
		result = helper90(nums, i+1, comb, result)
		comb = comb[:len(comb)-1]

		temp := i
		for i < len(nums) && nums[i] == nums[temp] {
			i++
		}

		result = helper90(nums, i, comb, result)
	}

	return result
}

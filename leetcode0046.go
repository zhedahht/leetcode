/*
LeetCode 46: https://leetcode.com/problems/permutations/
*/

package leetcode

func permute(nums []int) [][]int {
	return helper46(nums, 0, make([][]int, 0))
}

func helper46(nums []int, i int, result [][]int) [][]int {
	if i == len(nums) {
		clone := make([]int, len(nums))
		copy(clone, nums)
		result = append(result, clone)
	} else {
		for j := i; j < len(nums); j++ {
			nums[i], nums[j] = nums[j], nums[i]
			result = helper46(nums, i+1, result)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return result
}

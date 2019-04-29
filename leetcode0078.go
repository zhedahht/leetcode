/*
LeetCode 78: https://leetcode.com/problems/subsets/
*/

package leetcode

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	return helper78(nums, 0, make([]int, 0), result)
}

func helper78(nums []int, i int, cur []int, result [][]int) [][]int {
	if i == len(nums) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		result = append(result, temp)
	} else {
		cur = append(cur, nums[i])
		result = helper78(nums, i+1, cur, result)
		cur = cur[:len(cur)-1]

		result = helper78(nums, i+1, cur, result)
	}

	return result
}

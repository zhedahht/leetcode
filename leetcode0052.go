/*
LeetCode 52: https://leetcode.com/problems/n-queens-ii/
*/

package leetcode

func totalNQueens(n int) int {
	nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}

	return helper52(nums, 0, 0)
}

func helper52(nums []int, i int, result int) int {
	if i == len(nums) {
		result++
	} else if i < len(nums) {
		for j := i; j < len(nums); j++ {
			nums[i], nums[j] = nums[j], nums[i]

			// NOTE: isValidBoard is in leetcode51.go
			if isValidBoard(nums, i) {
				result = helper52(nums, i+1, result)
			}

			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return result
}

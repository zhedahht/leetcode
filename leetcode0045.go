/*
LeetCode 45: https://leetcode.com/problems/jump-game-ii/
*/

package leetcode

func jump(nums []int) int {
	left, right, steps := 0, 0, 0
	for right < len(nums)-1 {
		max := right + 1
		for i := left; i <= right; i++ {
			if max < nums[i]+i {
				max = nums[i] + i
			}
		}

		left = right + 1
		right = max
		steps++
	}

	return steps
}

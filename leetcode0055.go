/*
LeetCode 55: https://leetcode.com/problems/jump-game/
*/

package leetcode

func canJump(nums []int) bool {
	left, right := 0, 0
	for right < len(nums)-1 {
		max := right
		for i := left; i <= right; i++ {
			if max < nums[i]+i {
				max = nums[i] + i
			}
		}

		if max <= right {
			return false
		}

		left = right + 1
		right = max
	}

	return true
}

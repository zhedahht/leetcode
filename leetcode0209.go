/*
LeetCode 209: https://leetcode.com/problems/minimum-size-subarray-sum/
*/

package leetcode

func minSubArrayLen(target int, nums []int) int {
	minLength := len(nums) + 1
	left, right := 0, 0
	sum := nums[0]
	for left <= right && right < len(nums) {
		if sum >= target {
			length := right - left + 1
			if length < minLength {
				minLength = length
			}

			sum -= nums[left]
			left++
		} else {
			right++
			if right < len(nums) {
				sum += nums[right]
			}
		}
	}
	if minLength > len(nums) {
		minLength = 0
	}
	return minLength
}

/*
LeetCode 189: https://leetcode.com/problems/rotate-array/
*/

package leetcode

// The function name should be rotate. Rename it to avoid conflicts.
func rotate198(nums []int, k int) {
	reverse := func(nums []int, start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start, end = start+1, end-1
		}
	}

	k = k % len(nums)
	reverse(nums, 0, len(nums)-k-1)
	reverse(nums, len(nums)-k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

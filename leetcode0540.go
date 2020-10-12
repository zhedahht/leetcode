/*
LeetCode 540: https://leetcode.com/problems/single-element-in-a-sorted-array/
*/

package leetcode

func singleNonDuplicate(nums []int) int {
	pair := len(nums) / 2
	start, end := 0, pair-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid*2] != nums[mid*2+1] {
			if mid == 0 || nums[2*(mid-1)] == nums[2*(mid-1)+1] {
				return nums[2*mid]
			}

			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return nums[len(nums)-1]
}

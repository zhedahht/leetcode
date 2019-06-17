/*
LeetCode 154: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/
*/

package leetcode

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right && nums[left] >= nums[right] {
		if left+1 == right {
			return nums[right]
		}

		mid := (left + right) / 2
		if nums[left] == nums[right] && nums[left] == nums[mid] {
			right--
		} else if nums[mid] >= nums[left] {
			left = mid
		} else {
			right = mid
		}
	}

	return nums[left]
}

/*
LeetCode 34: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
*/

package leetcode

func searchRange(nums []int, target int) []int {
	searchFirstOrLast := func(nums []int, target int, isFirst bool) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := (left + right) / 2
			if nums[mid] == target {
				if isFirst {
					if mid == 0 || nums[mid-1] != target {
						return mid
					}

					right = mid - 1
				} else {
					if mid == len(nums)-1 || nums[mid+1] != target {
						return mid
					}

					left = mid + 1
				}
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		return -1
	}

	result := make([]int, 0)
	result = append(result, searchFirstOrLast(nums, target, true))
	result = append(result, searchFirstOrLast(nums, target, false))
	return result
}

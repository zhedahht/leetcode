/*
https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
*/

package leetcode

// NOTE: The name should be search. Rename it to avoid conflicts with leetcode0033.go
func search81(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	binarySearch := func(nums []int, start, end, target int) bool {
		for start <= end {
			mid := (start + end) / 2
			if nums[mid] == target {
				return true
			}

			if nums[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

		return false
	}

	left, right, smallest := 0, len(nums)-1, 0
	for left < right && nums[left] >= nums[right] {
		if right == left+1 {
			smallest = right
			break
		}

		mid := (left + right) / 2
		if nums[left] == nums[right] && nums[left] == nums[mid] {
			smallest, left, right = right, left+1, right-1
			continue
		}

		if nums[mid] >= nums[left] && nums[mid] >= nums[right] {
			left = mid
		} else {
			right = mid
		}
	}

	if target >= nums[0] {
		end := smallest - 1
		if smallest == 0 {
			end = len(nums) - 1
		}

		return binarySearch(nums, 0, end, target)
	}

	return binarySearch(nums, smallest, len(nums)-1, target)
}

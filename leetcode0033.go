/*
LeetCode 33: https://leetcode.com/problems/search-in-rotated-sorted-array/
*/

package leetcode

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	binarySearch := func(nums []int, start, end, target int) int {
		for start <= end {
			mid := (start + end) / 2
			if nums[mid] == target {
				return mid
			}

			if nums[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

		return -1
	}

	left, right, smallest := 0, len(nums)-1, 0
	for nums[left] > nums[right] {
		if right == left+1 {
			smallest = right
			break
		}

		mid := (left + right) / 2
		if nums[mid] > nums[left] {
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

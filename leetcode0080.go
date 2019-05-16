/*
LeetCode 80: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
*/

package leetcode

// NOTE: The name should be removeDuplicates. Rename it to avoid conflict with others.
func removeDuplicates80(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	i, j, count, first := 1, 1, 1, nums[0]
	for j < len(nums) {
		if nums[j] != first {
			nums[i] = nums[j]
			i++
			count, first = 1, nums[j]
		} else if count < 2 {
			nums[i] = nums[j]
			i++
			count++
		}

		j++
	}

	return i
}

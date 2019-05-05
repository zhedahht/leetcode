/*
LeetCode 26: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
*/

package leetcode

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for j < len(nums) {
		nums[i] = nums[j]
		i++

		k := j
		for j < len(nums) && nums[k] == nums[j] {
			j++
		}
	}

	return i
}

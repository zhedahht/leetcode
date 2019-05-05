/*
LeetCode 27: https://leetcode.com/problems/remove-element/
*/

package leetcode

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}

		j++
	}

	return i
}

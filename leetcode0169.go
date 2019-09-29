/*
LeetCode 169: https://leetcode.com/problems/majority-element/
*/

package leetcode

func majorityElement(nums []int) int {
	var result, count int
	for i := range nums {
		if nums[i] == result {
			count++
		} else if count > 0 {
			count--
		} else {
			result, count = nums[i], 1
		}
	}

	return result
}

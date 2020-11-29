/*
LeetCode 136: https://leetcode.com/problems/single-number/
*/

package leetcode

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}

/*
LeetCode 238: https://leetcode.com/problems/product-of-array-except-self/
*/

package leetcode

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	product := 1
	for i := range nums {
		result[i] = product
		product *= nums[i]
	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= product
		product *= nums[i]
	}

	return result
}

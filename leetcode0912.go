/*
LeetCode 912: https://leetcode.com/problems/sort-an-array/
*/

package leetcode

import "math/rand"

func sortArray(nums []int) []int {
	helper912(nums, 0, len(nums)-1)
	return nums
}

func helper912(nums []int, left, right int) {
	partition := func(nums []int, left, right int) int {
		r := rand.Intn(right-left) + left
		nums[r], nums[right] = nums[right], nums[r]

		less := left - 1
		for i := left; i < right; i++ {
			if nums[i] < nums[right] {
				less++
				nums[less], nums[i] = nums[i], nums[less]
			}
		}

		less++
		nums[less], nums[right] = nums[right], nums[less]
		return less
	}

	if left >= right {
		return
	}

	p := partition(nums, left, right)
	helper912(nums, left, p-1)
	helper912(nums, p+1, right)
}

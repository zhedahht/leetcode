/*
LeetCode 41: https://leetcode.com/problems/first-missing-positive/
*/

package leetcode

func firstMissingPositive(nums []int) int {
	swap := func(nums []int, i int) {
		num := nums[i]
		nums[i] = 0
		for num >= 1 && num <= len(nums) && num != i+1 {
			temp := nums[num-1]
			nums[num-1] = num
			i = num - 1
			num = temp
		}

		if num == i+1 {
			nums[i] = num
		}
	}

	for i, num := range nums {
		if num < 1 || num > len(nums) {
			nums[i] = 0
		} else {
			swap(nums, i)
		}
	}

	for i, num := range nums {
		if num == 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}

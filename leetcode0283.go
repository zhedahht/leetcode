/*
LeetCode 283: https://leetcode.com/problems/move-zeroes/description/
*/

package leetcode

func moveZeroes(nums []int) {
	count := 0
	for i, num := range nums {
		if num != 0 {
			if i != count {
				nums[count] = nums[i]
			}

			count++
		}
	}

	for ; count < len(nums); count++ {
		nums[count] = 0
	}
}

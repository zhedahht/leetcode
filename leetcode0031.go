/*
LeetCode 31: https://leetcode.com/problems/next-permutation/
*/

package leetcode

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	j, k := i+1, len(nums)-1
	for j < k {
		nums[j], nums[k] = nums[k], nums[j]
		j, k = j+1, k-1
	}

	if i >= 0 {
		j = i + 1
		for nums[j] <= nums[i] {
			j++
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}

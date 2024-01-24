/*
LeetCode 698: https://leetcode.com/problems/partition-to-k-equal-sum-subsets/description/
*/

package leetcode

func canPartitionKSubsets(nums []int, k int) bool {
	maxNum, total := 0, 0
	for _, num := range nums {
		maxNum = max(maxNum, num)
		total += num
	}

	if total%k != 0 {
		return false
	}

	if maxNum > total/k {
		return false
	}

	sums := make([]int, k)
	return putNumsWithTargetSum(nums, 0, sums, total/k)
}

func putNumsWithTargetSum(nums []int, i int, sums []int, target int) bool {
	if i == len(nums) {
		return true
	}

	tried := make(map[int]bool)
	for _, sum := range sums {
		tried[sum] = false
	}

	for j, sum := range sums {
		if !tried[sum] && nums[i]+sum <= target {
			tried[sum] = true

			sums[j] += nums[i]
			if putNumsWithTargetSum(nums, i+1, sums, target) {
				return true
			}

			sums[j] -= nums[i]
		}
	}

	return false
}

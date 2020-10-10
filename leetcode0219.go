/*
LeetCode 219: https://leetcode.com/problems/contains-duplicate-ii/
*/

package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)
	i := 0
	for ; i < len(nums) && i <= k; i++ {
		if m[nums[i]] {
			return true
		}

		m[nums[i]] = true
	}

	for ; i < len(nums); i++ {
		m[nums[i-k-1]] = false
		if m[nums[i]] {
			return true
		}

		m[nums[i]] = true
	}

	return false
}

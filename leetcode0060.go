/*
LeetCode 60: https://leetcode.com/problems/permutation-sequence/
*/

package leetcode

func getPermutation(n int, k int) string {
	nums := make([]byte, n)
	for i := 0; i < n; i++ {
		nums[i] = byte(i) + '1'
	}

	for i := 2; i <= k; i++ {
		nextPermutation60(nums)
	}

	return string(nums)
}

// NOTE: The algorithm is same as LeetCode 31
func nextPermutation60(nums []byte) {
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

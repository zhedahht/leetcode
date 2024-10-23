/*
LeetCode 487: https://leetcode.com/problems/max-consecutive-ones-ii/description/
*/

package leetcode

// The name should be findMaxConsecutiveOnes. It's renamed to avoid conflicts.
func findMaxConsecutiveOnes487(nums []int) int {
	cur, prev, result := 0, 0, 0
	has0 := false
	for i, num := range nums {
		if num == 0 {
			has0 = true
			if i >= 1 && nums[i-1] == 0 {
				prev = 0
			} else {
				prev = cur
				cur = 0
			}
			result = max(result, prev+1)
		} else {
			cur++
			if has0 {
				result = max(result, cur+prev+1)
			} else {
				result = max(result, cur)
			}
		}
	}

	return max(result, 1)
}

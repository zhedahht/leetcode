/*
LeetCode 485: https://leetcode.com/problems/max-consecutive-ones/description/
*/

package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	start := -1
	for i, num := range nums {
		if num == 0 {
			start = -1
		} else {
			if start == -1 {
				start = i
			}

			result = max(result, i-start+1)
		}
	}

	return result
}

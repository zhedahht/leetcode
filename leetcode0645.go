/*
LeetCode 643: https://leetcode.com/problems/maximum-average-subarray-i/description/
*/

package leetcode

func findErrorNums(nums []int) []int {
	result := make([]int, 2)
	for i, v := range nums {
		idx, val := i, v
		for val != idx+1 {
			newVal := nums[val-1]
			nums[val-1] = val
			if val == newVal {
				result[0] = val
				nums[idx] = val
				break
			} else {
				nums[val-1] = val
				val = newVal
			}
		}

		if nums[idx] != val {
			nums[idx] = val
		}
	}

	for idx, val := range nums {
		if val != idx+1 {
			result[1] = idx + 1
		}
	}

	return result
}

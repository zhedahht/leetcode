/*
LeetCode 491: https://leetcode.com/problems/non-decreasing-subsequences/
*/

package leetcode

import "fmt"

func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	strResult := make(map[string]bool)
	return helper491(nums, 0, make([]int, 0), result, strResult)
}

func helper491(nums []int, index int, cur []int, result [][]int, strResult map[string]bool) [][]int {
	if index == len(nums) {
		if len(cur) > 1 {
			str := fmt.Sprint("%v", cur)
			if _, exists := strResult[str]; !exists {
				strResult[str] = true

				clone := make([]int, len(cur))
				copy(clone, cur)
				result = append(result, clone)
			}
		}
	} else {
		result = helper491(nums, index+1, cur, result, strResult)

		if len(cur) == 0 || nums[index] >= cur[len(cur)-1] {
			cur := append(cur, nums[index])
			result = helper491(nums, index+1, cur, result, strResult)
			cur = cur[:len(cur)-1]
		}
	}

	return result
}

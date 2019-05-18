/*
LeetCode 51: https://leetcode.com/problems/n-queens/
*/

package leetcode

import (
	"fmt"
	"math"
)

func solveNQueens(n int) [][]string {
	nums := make([]int, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}

	return helper51(nums, 0, make([][]string, 0))
}

func isValidBoard(nums []int, last int) bool {
	for i := 0; i <= last; i++ {
		for j := i + 1; j <= last; j++ {
			if j-i == int(math.Abs(float64(nums[i]-nums[j]))) {
				return false
			}
		}
	}

	return true
}

func generateBoard(nums []int) []string {
	result := make([]string, 0)
	for _, num := range nums {
		var str string
		for i := 0; i < len(nums); i++ {
			if i == num {
				str = fmt.Sprintf("%sQ", str)
			} else {
				str = fmt.Sprintf("%s.", str)
			}
		}

		result = append(result, str)
	}

	return result
}

func helper51(nums []int, i int, result [][]string) [][]string {
	if i == len(nums) {
		result = append(result, generateBoard(nums))
	} else if i < len(nums) {
		for j := i; j < len(nums); j++ {
			nums[i], nums[j] = nums[j], nums[i]
			if isValidBoard(nums, i) {
				result = helper51(nums, i+1, result)
			}

			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return result
}

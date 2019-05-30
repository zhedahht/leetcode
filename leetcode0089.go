/*
LeetCode 89: https://leetcode.com/problems/gray-code/
*/

package leetcode

func grayCode(n int) []int {
	return helper89(n)
}

func helper89(n int) []int {
	if n == 0 {
		return make([]int, 1)
	}

	result := helper89(n - 1)
	for i := len(result) - 1; i >= 0; i-- {
		result = append(result, (1<<uint((n-1)))|result[i])
	}

	return result
}

/*
LeetCode 216: https://leetcode.com/problems/combination-sum-iii/
*/

package leetcode

func helper216(k, idx, n int, com []int, result [][]int) [][]int {
	if len(com) == k && n == 0 {
		c := make([]int, len(com))
		copy(c, com)
		result = append(result, c)
	} else if idx <= 9 && len(com) < k && n > 0 {
		result = helper216(k, idx+1, n, com, result)

		com = append(com, idx)
		result = helper216(k, idx+1, n-idx, com, result)
		com = com[:len(com)-1]
	}

	return result
}

func combinationSum3(k int, n int) [][]int {
	return helper216(k, 1, n, make([]int, 0), make([][]int, 0))
}

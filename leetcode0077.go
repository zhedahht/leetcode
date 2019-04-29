/*
LeetCode 77: https://leetcode.com/problems/combinations/
*/

package leetcode

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	cur := make([]int, 0)
	return helper77(n, k, 1, cur, result)
}

func helper77(n, k, i int, cur []int, result [][]int) [][]int {
	if k == 0 {
		temp := make([]int, len(cur))
		copy(temp, cur)
		result = append(result, temp)
	} else if i <= n {
		cur = append(cur, i)
		result = helper77(n, k-1, i+1, cur, result)
		cur = cur[:len(cur)-1]

		result = helper77(n, k, i+1, cur, result)
	}

	return result
}

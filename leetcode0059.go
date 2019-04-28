/*
LeetCode 59: https://leetcode.com/problems/spiral-matrix-ii/
*/

package leetcode

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	num := 1
	for i := 0; 2*i < n; i++ {
		if 2*i == n-1 {
			result[i][i] = num
		}

		for j := i; j < n-1-i; j++ {
			result[i][j] = num
			num++
		}

		for j := i; j < n-1-i; j++ {
			result[j][n-1-i] = num
			num++
		}

		for j := n - 1 - i; j > i; j-- {
			result[n-1-i][j] = num
			num++
		}

		for j := n - 1 - i; j > i; j-- {
			result[j][i] = num
			num++
		}
	}

	return result
}

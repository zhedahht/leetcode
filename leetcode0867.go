/*
LeetCode 867: https://leetcode.com/problems/transpose-matrix/
*/

package leetcode

func transpose(A [][]int) [][]int {
	result := make([][]int, len(A[0]))
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(A))
	}

	for i, row := range A {
		for j, num := range row {
			result[j][i] = num
		}
	}

	return result
}

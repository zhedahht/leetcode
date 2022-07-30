/*
LeetCode 304: https://leetcode.com/problems/range-sum-query-2d-immutable/
*/

package leetcode

type NumMatrix struct {
	sums [][]int
}

// Constructor0304 should be Constructor. It's renamed to avoid conflicts.
func Constructor0304(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix)+1)
	for i := range sums {
		sums[i] = make([]int, len(matrix[0])+1)
	}

	for j := 0; j < len(sums[0]); j++ {
		sums[0][j] = 0
	}

	for i := range matrix {
		sums[i+1][0] = 0
		sum := 0
		for j := range matrix[i] {
			sum += matrix[i][j]
			sums[i+1][j+1] = sums[i][j+1] + sum
		}
	}
	return NumMatrix{sums: sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2+1][col2+1] -
		this.sums[row2+1][col1] -
		this.sums[row1][col2+1] +
		this.sums[row1][col1]
}

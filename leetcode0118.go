/*
LeetCode 118: https://leetcode.com/problems/pascals-triangle/
*/

package leetcode

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	if numRows <= 0 {
		return result
	}

	result = append(result, []int{1})
	for i := 1; i < numRows; i++ {
		row := []int{1}
		for j := 1; j < i; j++ {
			row = append(row, result[i-1][j-1]+result[i-1][j])
		}

		row = append(row, 1)
		result = append(result, row)
	}

	return result
}

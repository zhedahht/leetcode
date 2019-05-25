/*
LeetCode 48: https://leetcode.com/problems/rotate-image/
*/

package leetcode

func rotate(matrix [][]int) {
	length := len(matrix)
	for i := 0; 2*i < length-1; i++ {
		for j := i; j < length-i-1; j++ {
			matrix[i][j], matrix[j][length-i-1], matrix[length-i-1][length-j-1], matrix[length-j-1][i] =
				matrix[length-j-1][i], matrix[i][j], matrix[j][length-i-1], matrix[length-i-1][length-j-1]
		}
	}
}

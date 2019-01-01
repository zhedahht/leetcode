/*
LeetCode 119: https://leetcode.com/problems/pascals-triangle-ii/
*/

package leetcode

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return make([]int, 0)
	}

	row := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		prev := 1
		for j := 1; j < i; j++ {
			num := prev + row[j]
			prev = row[j]
			row[j] = num
		}

		row[i] = 1
	}

	return row
}

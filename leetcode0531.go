/*
LeetCode 531: https://leetcode.com/problems/lonely-pixel-i/description/
*/

package leetcode

func findLonelyPixel(picture [][]byte) int {
	rows, cols := len(picture), len(picture[0])
	rowCounts, colCounts := make([]int, rows), make([]int, cols)
	for i, row := range picture {
		for j, val := range row {
			if val == 'B' {
				rowCounts[i]++
				colCounts[j]++
			}
		}
	}

	result := 0
	for i, row := range picture {
		for j, val := range row {
			if val == 'B' && rowCounts[i] == 1 && colCounts[j] == 1 {
				result++
			}
		}
	}

	return result
}

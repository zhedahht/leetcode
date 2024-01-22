/*
LeetCode 688: https://leetcode.com/problems/knight-probability-in-chessboard/description/
*/

package leetcode

func initMatrix(n int) [][]float64 {
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
	}

	return matrix
}

func knightProbability(n int, k int, row int, column int) float64 {
	dp1 := initMatrix(n)
	dp1[row][column] = 1
	dirs := [][]int{{1, 2}, {2, 1}, {-1, 2}, {1, -2}, {2, -1}, {-2, 1}, {-1, -2}, {-2, -1}}
	probability := 1.0
	for i := 0; i < k; i++ {
		dp2 := initMatrix(n)
		probability = 0.0
		for r := range dp1 {
			for c := range dp1[r] {
				for _, dir := range dirs {
					next := [2]int{r + dir[0], c + dir[1]}
					if next[0] >= 0 && next[0] < n && next[1] >= 0 && next[1] < n {
						dp2[next[0]][next[1]] += dp1[r][c] / 8.0
						probability += dp1[r][c] / 8.0
					}
				}
			}
		}

		dp1 = dp2
	}

	return probability
}

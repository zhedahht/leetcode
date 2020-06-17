package leetcode

func updateMatrix(matrix [][]int) [][]int {
	dist := make([][]int, len(matrix))
	visited := make([][]bool, len(matrix))
	queue := make([][]int, 0)

	for i, row := range matrix {
		dist[i] = make([]int, len(row))
		visited[i] = make([]bool, len(row))
		for j, v := range row {
			if v == 0 {
				queue = append(queue, []int{i, j})
				visited[i][j] = true
			}
		}
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			r := dir[0] + pos[0]
			c := dir[1] + pos[1]
			if r >= 0 && r < len(matrix) &&
				c >= 0 && c < len(matrix[0]) &&
				!visited[r][c] {
				visited[r][c] = true
				queue = append(queue, []int{r, c})
				dist[r][c] = dist[pos[0]][pos[1]] + 1
			}
		}
	}

	return dist
}

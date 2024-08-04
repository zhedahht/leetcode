/*
LeetCode 417: https://leetcode.com/problems/pacific-atlantic-water-flow/description/
*/

package leetcode

func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	flowed := make([][]int, rows)
	for i := 0; i < len(heights); i++ {
		flowed[i] = make([]int, cols)
	}

	queue := make([][]int, 0)
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	for i := 0; i < cols; i++ {
		queue = append(queue, []int{0, i})
		visited[0][i] = true
	}

	for i := 1; i < rows; i++ {
		queue = append(queue, []int{i, 0})
		visited[i][0] = true
	}

	reverseFlow(heights, queue, visited, flowed)

	queue = make([][]int, 0)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	for i := 0; i < cols; i++ {
		queue = append(queue, []int{rows - 1, i})
		visited[rows-1][i] = true
	}

	for i := 0; i < rows-1; i++ {
		queue = append(queue, []int{i, cols - 1})
		visited[i][cols-1] = true
	}

	reverseFlow(heights, queue, visited, flowed)

	result := make([][]int, 0)
	for i := 0; i < len(flowed); i++ {
		for j := 0; j < len(flowed[0]); j++ {
			if flowed[i][j] == 2 {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

func reverseFlow(heights [][]int, queue [][]int, visited [][]bool, flowed [][]int) {
	rows, cols := len(heights), len(heights[0])
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		flowed[pos[0]][pos[1]]++

		for _, dir := range dirs {
			next := []int{pos[0] + dir[0], pos[1] + dir[1]}
			if next[0] >= 0 && next[0] < rows && next[1] >= 0 && next[1] < cols &&
				!visited[next[0]][next[1]] &&
				heights[next[0]][next[1]] >= heights[pos[0]][pos[1]] {
				queue = append(queue, next)
				visited[next[0]][next[1]] = true
			}
		}
	}
}

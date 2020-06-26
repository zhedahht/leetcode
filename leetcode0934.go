/*
LeetCode 934: https://leetcode.com/problems/shortest-bridge/
*/

package leetcode

func shortestBridge(A [][]int) int {
	visited := make([][]bool, len(A))
	for i := range visited {
		visited[i] = make([]bool, len(A[0]))
	}

	queue1, queue2 := make([][]int, 0), make([][]int, 0)
	island := make([][]int, 0)
	for i := 0; i < len(A) && len(island) == 0; i++ {
		for j := 0; j < len(A[0]) && len(island) == 0; j++ {
			if A[i][j] == 1 {
				queue1 = append(queue1, []int{i, j})
				visited[i][j] = true
				island = append(island, []int{i, j})
				break
			}
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue1) > 0 {
		r, c := queue1[0][0], queue1[0][1]
		queue1 = queue1[1:]

		for _, dir := range dirs {
			nextR, nextC := r+dir[0], c+dir[1]
			if nextR >= 0 && nextR < len(A) &&
				nextC >= 0 && nextC < len(A[0]) &&
				!visited[nextR][nextC] && A[nextR][nextC] == 1 {
				queue1 = append(queue1, []int{nextR, nextC})
				island = append(island, []int{nextR, nextC})

				visited[nextR][nextC] = true
			}
		}
	}

	queue1 = island
	dist := 0
	for len(queue1) > 0 {
		r, c := queue1[0][0], queue1[0][1]
		queue1 = queue1[1:]
		visited[r][c] = true

		for _, dir := range dirs {
			nextR, nextC := r+dir[0], c+dir[1]
			if nextR >= 0 && nextR < len(A) &&
				nextC >= 0 && nextC < len(A[0]) &&
				!visited[nextR][nextC] {
				if A[nextR][nextC] == 1 {
					return dist
				}

				queue2 = append(queue2, []int{nextR, nextC})
				visited[nextR][nextC] = true
			}
		}

		if len(queue1) == 0 {
			dist++
			queue1 = queue2
			queue2 = make([][]int, 0)
		}
	}

	return -1
}

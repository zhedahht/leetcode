/*
LeetCode 886: https://leetcode.com/problems/possible-bipartition/
*/

package leetcode

func possibleBipartition(N int, dislikes [][]int) bool {
	numToPrev, numToNext := make(map[int][]int), make(map[int][]int)
	for i := 1; i <= N; i++ {
		numToPrev[i], numToNext[i] = make([]int, 0), make([]int, 0)
	}

	for _, dislike := range dislikes {
		prev, next := dislike[0], dislike[1]
		numToPrev[next] = append(numToPrev[next], prev)
		numToNext[prev] = append(numToNext[prev], next)
	}

	visited := make(map[int]int)
	for i := 1; i <= N; i++ {
		if _, exists := visited[i]; !exists {
			if !helper886(i, 0, numToPrev, numToNext, visited) {
				return false
			}
		}
	}

	return true
}

func helper886(num, level int, numToPrev, numToNext map[int][]int, visited map[int]int) bool {
	visitNeighbor := func(neighbor, level int, visited map[int]int) bool {
		if l, exists := visited[neighbor]; exists {
			if l%2 != level%2 {
				return false
			}
		} else {
			visited[neighbor] = level
			if !helper886(neighbor, level, numToPrev, numToNext, visited) {
				return false
			}
		}

		return true
	}

	for _, prev := range numToPrev[num] {
		if !visitNeighbor(prev, level+1, visited) {
			return false
		}
	}

	for _, next := range numToNext[num] {
		if !visitNeighbor(next, level+1, visited) {
			return false
		}
	}

	return true
}

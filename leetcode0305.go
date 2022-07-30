/*
LeetCode 305: https://leetcode.com/problems/number-of-islands-ii/
*/

package leetcode

func numIslands2(m int, n int, positions [][]int) []int {
	getIndex := func(i, j int) int {
		return i*n + j
	}

	result := make([]int, len(positions))
	fathers := make(map[int]int)
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for i, pos := range positions {
		prev := 0
		if i > 0 {
			prev = result[i-1]
		}

		index := getIndex(pos[0], pos[1])
		if _, exists := fathers[index]; exists {
			result[i] = prev
			continue
		}

		fathers[index] = index
		islands := make(map[int]bool)
		merge := 0
		for _, dir := range dirs {
			nb := []int{pos[0] + dir[0], pos[1] + dir[1]}
			if nb[0] < 0 || nb[0] >= m || nb[1] < 0 || nb[1] >= n {
				continue
			}

			indexNB := getIndex(nb[0], nb[1])
			if _, exists := fathers[indexNB]; !exists {
				continue
			}

			island := findFather0305(fathers, indexNB)
			if islands[island] {
				continue
			}

			islands[island] = true
			union0305(fathers, indexNB, index)
			merge++
		}

		result[i] = prev + 1 - merge
	}

	return result
}

func findFather0305(fathers map[int]int, i int) int {
	if fathers[i] != i {
		fathers[i] = findFather0305(fathers, fathers[i])
	}
	return fathers[i]
}

func union0305(fathers map[int]int, i, j int) {
	fatherI := findFather0305(fathers, i)
	fatherJ := findFather0305(fathers, j)
	fathers[fatherJ] = fatherI
}

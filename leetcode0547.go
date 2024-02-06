/*
LeetCode 547: https://leetcode.com/problems/number-of-provinces/description/
*/

package leetcode

func findCircleNum(isConnected [][]int) int {
	parents := make(map[int]int)
	for i := 0; i < len(isConnected); i++ {
		parents[i] = i
	}

	for i := range isConnected {
		for j := 0; j < i; j++ {
			if isConnected[i][j] == 1 {
				union547(parents, i, j)
			}
		}
	}

	circles := make(map[int]bool)
	for i := 0; i < len(parents); i++ {
		parent := getParent547(parents, i)
		circles[parent] = true
	}

	return len(circles)
}

func union547(parents map[int]int, i, j int) {
	parent1 := getParent547(parents, i)
	parent2 := getParent547(parents, j)
	parents[parent1] = parents[parent2]
}

func getParent547(parents map[int]int, i int) int {
	parent := parents[i]
	if parent != parents[parent] {
		parent = getParent547(parents, parent)
		parents[i] = parent
	}

	return parent
}

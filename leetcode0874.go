/*
LeetCode 874: https://leetcode.com/problems/walking-robot-simulation/
*/

package leetcode

import "math"

func robotSim(commands []int, obstacles [][]int) int {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIndex := 0
	obs := make(map[[2]int]bool)
	for _, ob := range obstacles {
		obs[[2]int{ob[0], ob[1]}] = true
	}

	pos := [2]int{0, 0}
	result := 0
	for _, command := range commands {
		if command == -1 {
			dirIndex = (dirIndex + 1) % 4
		} else if command == -2 {
			dirIndex = (dirIndex + 3) % 4
		} else {
			for i := 0; i < command; i++ {
				next := [2]int{pos[0] + dirs[dirIndex][0], pos[1] + dirs[dirIndex][1]}
				if _, exists := obs[next]; exists {
					break
				}

				pos = next
				dist := pos[0]*pos[0] + pos[1]*pos[1]
				result = int(math.Max(float64(result), float64(dist)))
			}
		}
	}

	return result
}

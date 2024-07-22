/*
LeetCode 401: https://leetcode.com/problems/binary-watch/description/
*/

package leetcode

import "fmt"

func readBinaryWatch(turnedOn int) []string {
	result := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if getBit1Count(i)+getBit1Count(j) == turnedOn {
				time := fmt.Sprintf("%d:%02d", i, j)
				result = append(result, time)
			}
		}
	}

	return result
}

func getBit1Count(num int) int {
	count := 0
	for num != 0 {
		count++
		num = num & (num - 1)
	}

	return count
}

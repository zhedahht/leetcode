/*
LeetCode 621: https://leetcode.com/problems/task-scheduler/
*/

package leetcode

func leastInterval(tasks []byte, n int) int {
	counts := make(map[byte]int)
	for _, task := range tasks {
		counts[task]++
	}

	max := 0
	for _, count := range counts {
		if max < count {
			max = count
		}
	}

	maxCount := 0
	for _, count := range counts {
		if count == max {
			maxCount++
		}
	}

	all, others := len(tasks), len(tasks)-max
	if others-maxCount+1 >= n*(max-1) {
		return all
	}

	return (n+1)*(max-1) + maxCount
}

/*
LeetCode 495: https://leetcode.com/problems/teemo-attacking/description/
*/

package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	start, end := -1, -1
	result := 0
	for _, t := range timeSeries {
		if t > end {
			result += (end - start)
			start, end = t, t+duration
		} else {
			end = t + duration
		}
	}

	result += (end - start)
	return result
}

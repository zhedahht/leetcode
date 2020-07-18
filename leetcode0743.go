/*
LeetCode 743: https://leetcode.com/problems/network-delay-time/
*/

package leetcode

func networkDelayTime(times [][]int, N int, K int) int {
	const maxInt = 0x7fffffff
	delay := make([]int, N+1)
	for i := range delay {
		delay[i] = maxInt
	}

	delay[K] = 0
	for i := 0; i < N-1; i++ {
		for _, t := range times {
			if delay[t[0]] < maxInt && delay[t[1]] > delay[t[0]]+t[2] {
				delay[t[1]] = delay[t[0]] + t[2]
			}
		}
	}

	result := 0
	for i := 1; i < len(delay); i++ {
		if delay[i] == maxInt {
			return -1
		}

		if delay[i] > result {
			result = delay[i]
		}
	}

	return result
}

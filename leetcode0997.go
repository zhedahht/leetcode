/*
LeetCode 997: https://leetcode.com/problems/find-the-town-judge/
*/

package leetcode

func findJudge(N int, trust [][]int) int {
	in := make([]int, N)
	out := make([]int, N)

	for _, pair := range trust {
		out[pair[0]-1]++
		in[pair[1]-1]++
	}

	for i := 0; i < N; i++ {
		if in[i] == N-1 && out[i] == 0 {
			return i + 1
		}
	}

	return -1
}

/*
LeetCode 646: https://leetcode.com/problems/maximum-length-of-pair-chain/description/
*/

package leetcode

import "sort"

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})

	lengths := make([]int, len(pairs))
	for i, pair := range pairs {
		lengths[i] = 1
		for j := 0; j < i; j++ {
			prev := pairs[j]
			if prev[1] < pair[0] && lengths[i] < lengths[j]+1 {
				lengths[i] = lengths[j] + 1
			}
		}
	}

	max := 0
	for _, length := range lengths {
		if max < length {
			max = length
		}
	}

	return max
}

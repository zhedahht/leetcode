/*
LeetCode 191: https://leetcode.com/problems/number-of-1-bits/
*/

package leetcode

func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		result++
		num = num & (num - 1)
	}

	return result
}

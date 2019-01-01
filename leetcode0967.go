/*
LeetCode: https://leetcode.com/problems/numbers-with-same-consecutive-differences/
*/

package leetcode

import "strconv"

func numsSameConsecDiff(N int, K int) []int {
	result := make([]int, 0)
	num := make([]byte, 0)
	start := 1
	if N == 1 {
		start = 0
	}

	for i := start; i <= 9; i++ {
		num = append(num, byte(i)+'0')
		numsSameConsecDiffCore(N, K, 0, byte(i)+'0', num, &result)
		num = num[0 : len(num)-1]
	}

	return result
}

func numsSameConsecDiffCore(N int, K int, index int, digit byte, num []byte, result *[]int) {
	if index == N-1 {
		val, _ := strconv.Atoi(string(num))
		*result = append(*result, val)
		return
	}

	if digit-byte(K) >= '0' {
		numsSameConsecDiffCore(N, K, index+1, digit-byte(K), append(num, digit-byte(K)), result)
	}

	if K != 0 && digit+byte(K) <= '9' {
		numsSameConsecDiffCore(N, K, index+1, digit+byte(K), append(num, digit+byte(K)), result)
	}
}

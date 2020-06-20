/*
LeetCode 985: https://leetcode.com/problems/sum-of-even-numbers-after-queries/
*/

package leetcode

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sumEven, sumOdd := 0, 0
	for _, v := range A {
		if v%2 == 0 {
			sumEven += v
		} else {
			sumOdd += v
		}
	}

	result := make([]int, len(queries))
	for i, q := range queries {
		idx, val := q[1], q[0]
		if val%2 == 0 {
			if A[idx]%2 == 0 {
				sumEven += val
			} else {
				sumOdd += val
			}
		} else {
			if A[idx]%2 == 0 {
				sumOdd += (A[idx] + val)
				sumEven -= A[idx]
			} else {
				sumOdd -= A[idx]
				sumEven += (A[idx] + val)
			}
		}

		result[i] = sumEven
		A[idx] += val
	}

	return result
}

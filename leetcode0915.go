/*
LeetCode 915: https://leetcode.com/problems/partition-array-into-disjoint-intervals/
*/

package leetcode

func partitionDisjoint(A []int) int {
	min, max := make([]int, len(A)), make([]int, len(A))

	max[0] = A[0]
	for i := 1; i < len(A); i++ {
		if A[i] > max[i-1] {
			max[i] = A[i]
		} else {
			max[i] = max[i-1]
		}
	}

	min[len(A)-1] = A[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] < min[i+1] {
			min[i] = A[i]
		} else {
			min[i] = min[i+1]
		}
	}

	for i := 0; i < len(A)-1; i++ {
		if max[i] <= min[i+1] {
			return i + 1
		}
	}

	return -1
}

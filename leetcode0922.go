/*
LeetCode 922: https://leetcode.com/problems/sort-array-by-parity-ii/
*/

package leetcode

func sortArrayByParityII(A []int) []int {
	i := 0
	j := 1
	for i < len(A) && j < len(A) {
		for i < len(A) && A[i]%2 == 0 {
			i += 2
		}

		for j < len(A) && A[j]%2 == 1 {
			j += 2
		}

		if i < len(A) && j < len(A) {
			A[i], A[j] = A[j], A[i]
			i += 2
			j += 2
		}
	}

	return A
}

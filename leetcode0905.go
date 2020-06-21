/*
LeetCode 905: https://leetcode.com/problems/sort-array-by-parity/
*/

package leetcode

func sortArrayByParity(A []int) []int {
	lastEven := -1
	for i := range A {
		if A[i]%2 == 0 {
			lastEven++
			A[lastEven], A[i] = A[i], A[lastEven]
		}
	}

	return A
}

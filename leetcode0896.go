/*
https://leetcode.com/problems/monotonic-array/
*/

package leetcode

func isMonotonic(A []int) bool {
	isInc, isDec := true, true
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			isDec = false
		} else if A[i] < A[i-1] {
			isInc = false
		}

		if !(isInc || isDec) {
			return false
		}
	}

	return true
}

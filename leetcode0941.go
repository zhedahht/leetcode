/*
LeetCode 941: https://leetcode.com/problems/valid-mountain-array/
*/

package leetcode

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	i := 1
	for ; i < len(A); i++ {
		if A[i] == A[i-1] {
			return false
		}

		if A[i] < A[i-1] {
			break
		}
	}

	if i == 1 || i == len(A) {
		return false
	}

	for ; i < len(A); i++ {
		if A[i] >= A[i-1] {
			return false
		}
	}

	return true
}

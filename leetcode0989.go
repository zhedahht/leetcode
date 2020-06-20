/*
LeetCode 989: https://leetcode.com/problems/add-to-array-form-of-integer/
*/

package leetcode

func addToArrayForm(A []int, K int) []int {
	toSlice := func(K int) []int {
		s := make([]int, 0)
		for K > 0 {
			s = append(s, K%10)
			K = K / 10
		}

		return s
	}

	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	s := toSlice(K)
	reverse(A)
	sum := make([]int, 0)
	i, carry := 0, 0
	for i < len(A) || i < len(s) {
		a := carry
		if i < len(A) {
			a += A[i]
		}

		if i < len(s) {
			a += s[i]
		}

		if a >= 10 {
			a -= 10
			carry = 1
		} else {
			carry = 0
		}

		sum = append(sum, a)
		i++
	}

	if carry > 0 {
		sum = append(sum, carry)
	}

	reverse(sum)
	return sum
}

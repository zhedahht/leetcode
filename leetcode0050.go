/*
LeetCode 50: https://leetcode.com/problems/powx-n/
*/

package leetcode

func myPow(x float64, n int) float64 {
	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	result := helper50(x, n)
	if isNegative {
		result = 1 / result
	}

	return result
}

func helper50(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	temp := helper50(x, n/2)
	if n%2 == 0 {
		return temp * temp
	}

	return x * temp * temp
}

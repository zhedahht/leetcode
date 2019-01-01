/*
LeetCode 633: https://leetcode.com/problems/sum-of-square-numbers/
*/

package leetcode

import "math"

func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}

	for i := 0; i <= int(math.Sqrt(float64(c)/2.0)); i++ {
		a := i * i
		b := c - a
		if int(math.Sqrt(float64(b)))*int(math.Sqrt(float64(b))) == b {
			return true
		}
	}

	return false
}

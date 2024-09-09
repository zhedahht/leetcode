/*
LeetCode 470: https://leetcode.com/problems/implement-rand10-using-rand7/description/
*/

package leetcode

import "math/rand"

// rand7 is added to avoid compilation errors
func rand7() int {
	return rand.Int()%7 + 1
}

func rand10() int {
	a, b := rand7(), rand7()
	num := (a-1)*7 + (b - 1)
	if num > 39 {
		return rand10()
	}

	return num%10 + 1
}

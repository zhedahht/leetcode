/*
LeetCode 640: https://leetcode.com/problems/solve-the-equation/
*/

package leetcode

import (
	"fmt"
	"strings"
)

func solveEquation(equation string) string {
	segs := strings.Split(equation, "=")
	x1, val1 := normalize(segs[0])
	x2, val2 := normalize(segs[1])

	x := x1 - x2
	val := val2 - val1

	if x == 0 && val == 0 {
		return "Infinite solutions"
	}

	if x == 0 {
		return "No solution"
	}

	return fmt.Sprintf("x=%d", val/x)
}

func normalize(seg string) (int, int) {
	x := 0
	val := 0
	index := 0
	for index < len(seg) {
		negative := false
		if seg[index] == '-' {
			negative = true
			index++
		} else if seg[index] == '+' {
			index++
		}

		num := 0
		hasDigits := false
		for index < len(seg) && seg[index] >= '0' && seg[index] <= '9' {
			hasDigits = true
			num = num*10 + int(seg[index]-'0')
			index++
		}

		if !hasDigits {
			num = 1
		}

		if negative {
			num = -num
		}

		if index < len(seg) && seg[index] == 'x' {
			x += num
			index++
		} else {
			val += num
		}
	}

	return x, val
}

/*
LeetCode 150: https://leetcode.com/problems/evaluate-reverse-polish-notation/
*/

package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	getTwoVals := func(vals []int) (int, int, []int) {
		le := len(vals)
		return vals[le-2], vals[le-1], vals[:le-2]
	}
	vals := make([]int, 0)
	for _, token := range tokens {
		var val1, val2 int
		switch token {
		case "+":
			val1, val2, vals = getTwoVals(vals)
			vals = append(vals, val1+val2)
		case "-":
			val1, val2, vals = getTwoVals(vals)
			vals = append(vals, val1-val2)
		case "*":
			val1, val2, vals = getTwoVals(vals)
			vals = append(vals, val1*val2)
		case "/":
			val1, val2, vals = getTwoVals(vals)
			vals = append(vals, val1/val2)
		default:
			val, _ := strconv.Atoi(token)
			vals = append(vals, val)
		}
	}
	return vals[0]
}

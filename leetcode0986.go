/*
LeetCode 986: https://leetcode.com/problems/interval-list-intersections/
*/

package leetcode

import "math"

func intervalIntersection(A [][]int, B [][]int) [][]int {
	intersect := func(a, b []int) []int {
		if a[0] > b[0] {
			a, b = b, a
		}

		if b[0] <= a[1] {
			return []int{b[0], int(math.Min(float64(a[1]), float64(b[1])))}
		}

		return nil
	}

	i, j := 0, 0
	result := make([][]int, 0)
	for i < len(A) && j < len(B) {
		in := intersect(A[i], B[j])
		if in != nil {
			result = append(result, in)
		}

		if A[i][1] <= B[j][1] {
			i++
		} else {
			j++
		}
	}

	return result
}

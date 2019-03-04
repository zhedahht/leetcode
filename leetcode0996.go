/*
LeetCode 996: https://leetcode.com/problems/number-of-squareful-arrays/
*/

package leetcode

import "math"

func numSquarefulPerms(A []int) int {
	return helper996(A, 0, 0)
}

func helper996(A []int, index, count int) int {
	if index == len(A) {
		return count + 1
	}

	visited := map[int]bool{}
	for i := index; i < len(A); i++ {
		if _, ok := visited[A[i]]; !ok {
			visited[A[i]] = true

			if index == 0 || isSquare(A[index-1]+A[i]) {
				A[index], A[i] = A[i], A[index]
				count = helper996(A, index+1, count)
				A[index], A[i] = A[i], A[index]
			}
		}
	}

	return count
}

func isSquare(num int) bool {
	root := int(math.Sqrt(float64(num)))
	return num == root*root
}

/*
LeetCode 593: https://leetcode.com/problems/valid-square/
*/

package leetcode

import "sort"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	dists := make([]int, 0)
	dists = append(dists, distSquare(p1, p2))
	dists = append(dists, distSquare(p1, p3))
	dists = append(dists, distSquare(p1, p4))
	dists = append(dists, distSquare(p2, p3))
	dists = append(dists, distSquare(p2, p4))
	dists = append(dists, distSquare(p3, p4))
	sort.Ints(dists)

	if dists[0] == 0 {
		return false
	}

	if dists[0] != dists[1] || dists[1] != dists[2] || dists[2] != dists[3] || dists[3] != dists[0] {
		return false
	}

	if dists[4] != dists[5] {
		return false
	}

	return true
}

func distSquare(p1, p2 []int) int {
	dist := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
	return dist
}

/*
LeetCode 658: https://leetcode.com/problems/find-k-closest-elements/description/
*/

package leetcode

import "math"

func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	for left+k < len(arr) {
		if math.Abs(float64(arr[left+k]-x)) < math.Abs(float64(arr[left]-x)) {
			left++
		} else if arr[left+k] == arr[left] {
			left++
		} else {
			break
		}
	}

	return arr[left : left+k]
}

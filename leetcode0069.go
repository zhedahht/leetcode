/*
LeetCode 69: https://leetcode.com/problems/sqrtx/
*/

package leetcode

func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square == x {
			return mid
		} else if square < x {
			if (mid+1)*(mid+1) > x {
				return mid
			}

			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

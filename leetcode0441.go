/*
LeetCode 441: https://leetcode.com/problems/arranging-coins/
*/

package leetcode

import "math"

func arrangeCoins(n int) int {
	left, right := 1, int(math.Sqrt(float64(2*n)))
	for left <= right {
		mid := left + (right-left)/2
		full := mid * (mid + 1) / 2
		if full == n {
			return mid
		}

		if full < n {
			nextFull := (mid + 1) * (mid + 2) / 2
			if nextFull > n {
				return mid
			}

			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return 0
}

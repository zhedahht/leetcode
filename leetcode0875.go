/*
LeetCode 875: https://leetcode.com/problems/koko-eating-bananas/
*/

package leetcode

import "math"

func minEatingSpeed(piles []int, H int) int {
	getTime := func(piles []int, speed int) int {
		result := 0
		for _, p := range piles {
			result += (p + speed - 1) / speed
		}

		return result
	}

	min, max := 1, 0
	for _, p := range piles {
		max = int(math.Max(float64(max), float64(p)))
	}

	for min <= max {
		mid := (min + max) / 2
		time := getTime(piles, mid)
		if time > H {
			min = mid + 1
		} else {
			if mid == 1 || getTime(piles, mid-1) > H {
				return mid
			}

			max = mid - 1
		}
	}

	return 0
}

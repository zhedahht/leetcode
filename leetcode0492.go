/*
LeetCode 492: https://leetcode.com/problems/construct-the-rectangle/description/
*/

package leetcode

import "math"

func constructRectangle(area int) []int {
	mid := int(math.Sqrt(float64(area)))
	for i := mid; i >= 1; i-- {
		j := area / i
		if i*j == area {
			return []int{j, i}
		}
	}

	return []int{area, 1}
}

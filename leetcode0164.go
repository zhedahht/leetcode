/*
LeetCode 164: https://leetcode.com/problems/maximum-gap/
*/

package leetcode

import "math"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	min := int((^(uint(0))) >> 1)
	max := 0
	for _, num := range nums {
		min = int(math.Min(float64(min), float64(num)))
		max = int(math.Max(float64(max), float64(num)))
	}

	total := max - min + 1
	bucketRange := float64(total) / float64(len(nums)-1)
	rangeMins := make([]int, len(nums)-1)
	for i, _ := range rangeMins {
		rangeMins[i] = int((^(uint(0))) >> 1)
	}

	rangeMaxs := make([]int, len(nums)-1)
	for i, _ := range rangeMaxs {
		rangeMaxs[i] = 0
	}

	for _, num := range nums {
		index := int(float64(num-min) / bucketRange)
		rangeMins[index] = int(math.Min(float64(rangeMins[index]), float64(num)))
		rangeMaxs[index] = int(math.Max(float64(rangeMaxs[index]), float64(num)))
	}

	gapMax := rangeMaxs[0] - rangeMins[0]
	prevMax := rangeMaxs[0]
	for i := 1; i < len(nums)-1; i++ {
		if rangeMaxs[i] >= rangeMins[i] {
			gap := rangeMaxs[i] - rangeMins[i]
			gapMax = int(math.Max(float64(gapMax), float64(gap)))

			gap = rangeMins[i] - prevMax
			gapMax = int(math.Max(float64(gapMax), float64(gap)))

			prevMax = rangeMaxs[i]
		}
	}

	return gapMax
}

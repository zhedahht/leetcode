/*
LeetCode 137: https://leetcode.com/problems/single-number-ii/
*/

package leetcode

func singleNumber137(nums []int) int {
	const intSize = 64
	digitSums := make([]int, intSize)
	for _, num := range nums {
		for i := intSize - 1; i >= 0; i-- {
			flag := 1 << i
			digit := (num & flag) >> i
			if digit != 0 {
				digitSums[intSize-1-i]++
			}
		}
	}

	unique := 0
	for i, s := range digitSums {
		if s%3 == 1 {
			unique |= 1 << (intSize - 1 - i)
		}
	}

	return unique
}

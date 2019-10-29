/*
LeetCode 869: https://leetcode.com/problems/reordered-power-of-2/
*/

package leetcode

import (
	"fmt"
	"sort"
)

func reorderedPowerOf2(N int) bool {
	getString := func(n int) string {
		str := fmt.Sprintf("%d", n)
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		return string(runes)
	}

	nums := make(map[string]bool)
	for i := uint(0); i < 31; i++ {
		n := 1 << i
		nums[getString(n)] = true
	}

	str := getString(N)
	return nums[str]
}

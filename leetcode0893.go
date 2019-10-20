/*
LeetCode 893: https://leetcode.com/problems/groups-of-special-equivalent-strings/
*/

package leetcode

import "sort"

func numSpecialEquivGroups(A []string) int {
	sortStr := func(str string) string {
		var odd, even []byte
		for i, b := range str {
			if i%2 == 0 {
				even = append(even, byte(b))
			} else {
				odd = append(odd, byte(b))
			}
		}

		sort.Slice(odd, func(i, j int) bool { return odd[i] < odd[j] })
		sort.Slice(even, func(i, j int) bool { return even[i] < even[j] })
		return string(even) + string(odd)
	}

	allSorted := make(map[string]bool)
	for _, str := range A {
		sorted := sortStr(str)
		allSorted[sorted] = true
	}

	return len(allSorted)
}

/*
LeetCode 791: https://leetcode.com/problems/custom-sort-string/
*/

package leetcode

import (
	"sort"
	"strings"
)

func customSortString(S string, T string) string {
	t := []byte(T)
	sort.Slice(t, func(i, j int) bool {
		idxI, idxJ := strings.IndexByte(S, t[i]), strings.IndexByte(S, t[j])
		return idxI < idxJ
	})

	return string(t)
}

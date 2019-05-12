/*
LeetCode 71: https://leetcode.com/problems/simplify-path/
*/

package leetcode

import (
	"strings"
)

func simplifyPath(path string) string {
	raw := strings.Split(path, "/")
	segs := make([]string, 0)
	for _, r := range raw {
		if r == ".." && len(segs) > 0 {
			segs = segs[:len(segs)-1]
		} else if r != ".." && r != "" && r != "." {
			segs = append(segs, r)
		}
	}

	var result string
	for _, seg := range segs {
		result = result + "/" + seg
	}

	if result == "" {
		result = "/"
	}

	return result
}

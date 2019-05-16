/*
LeetCode 38: https://leetcode.com/problems/count-and-say/
*/

package leetcode

import (
	"fmt"
)

func countAndSay(n int) string {
	result := "1"
	for i := 2; i <= n; i++ {
		temp := ""
		for j := 0; j < len(result); {
			k := j
			for k < len(result) && result[k] == result[j] {
				k++
			}

			temp = fmt.Sprintf("%s%d%c", temp, k-j, result[j])
			j = k
		}

		result = temp
	}

	return result
}

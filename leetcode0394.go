/*
LeetCode 394: https://leetcode.com/problems/decode-string/description/
*/

package leetcode

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
	strs := make([]string, 0)
	nums := make([]int, 0)
	for i := 0; i < len(s); {
		fmt.Println(s, i)
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}

			nums = append(nums, num)
			continue
		}

		if s[i] == '[' {
			i++
			segBuilder := strings.Builder{}
			for s[i] >= 'a' && s[i] <= 'z' {
				segBuilder.WriteByte(s[i])
				i++
			}

			strs = pushStringIntoStack(strs, segBuilder.String(), false)
			continue
		}

		if s[i] == ']' {
			i++
			num := nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			str := strs[len(strs)-1]
			strs = strs[:len(strs)-1]
			segBuilder := strings.Builder{}
			for i := 0; i < num; i++ {
				segBuilder.WriteString(str)
			}

			strs = pushStringIntoStack(strs, segBuilder.String(), true)
			continue
		}

		segBuilder := strings.Builder{}
		for i < len(s) && s[i] >= 'a' && s[i] <= 'z' {
			segBuilder.WriteByte(s[i])
			i++
		}
		strs = pushStringIntoStack(strs, segBuilder.String(), true)
	}

	if len(strs) == 0 {
		return ""
	}
	return strs[0]
}

func pushStringIntoStack(strs []string, str string, mergeWithPrev bool) []string {
	if mergeWithPrev {
		if len(strs) == 0 {
			strs = []string{str}
		} else {
			strs[len(strs)-1] = fmt.Sprintf("%s%s", strs[len(strs)-1], str)
		}
	} else {
		strs = append(strs, str)
	}

	return strs
}

/*
LeetCode 301: https://leetcode.com/problems/remove-invalid-parentheses/
*/

package leetcode

func removeInvalidParentheses(s string) []string {
	isValid := func(str string) bool {
		count, valid := 0, true
		for _, r := range str {
			if r == '(' {
				count++
			} else if r == ')' {
				if count <= 0 {
					valid = false
					break
				} else {
					count--
				}
			}
		}

		return count == 0 && valid
	}

	set1, set2 := make(map[string]bool), make(map[string]bool)
	set1[s] = true
	result := make([]string, 0)
	for len(set1) > 0 {
		for str := range set1 {
			if isValid(str) {
				result = append(result, str)
			} else if len(result) == 0 {
				for i := 0; i < len(str); i++ {
					strLeft := ""
					if i > 0 {
						strLeft = string(str[:i])
					}

					strRight := ""
					if i < len(str)-1 {
						strRight = string(str[i+1:])
					}

					set2[strLeft+strRight] = true
				}
			}
		}

		if len(result) > 0 {
			break
		}

		set1 = set2
		set2 = make(map[string]bool)
	}

	return result
}

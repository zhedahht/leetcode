/*
LeetCode https://leetcode.com/problems/remove-comments/
*/

package leetcode

func removeComments(source []string) []string {
	state := 0
	result := make([]string, 0)
	line := make([]byte, 0)
	for _, str := range source {
		for i := 0; i < len(str); {
			switch state {
			case 0:
				if str[i] == '/' && i < len(str)-1 {
					if str[i+1] == '/' {
						state = 1
						i += 2
					} else if str[i+1] == '*' {
						state = 2
						i += 2
					} else {
						line = append(line, str[i])
						i++
					}
				} else {
					line = append(line, str[i])
					i++
				}
			case 1:
				i++
			case 2:
				if str[i] == '*' && i < len(str)-1 && str[i+1] == '/' {
					state = 0
					i += 2
				} else {
					i++
				}
			}
		}

		if state == 1 {
			state = 0
		}

		if state == 0 && len(line) > 0 {
			result = append(result, string(line))
			line = make([]byte, 0)
		}
	}

	return result
}

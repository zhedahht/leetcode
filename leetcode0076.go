/*
LeetCode 76: https://leetcode.com/problems/minimum-window-substring/
*/

package leetcode

func minWindow(s string, t string) string {
	counts := make(map[byte]int)
	for _, ch := range []byte(t) {
		if _, exists := counts[ch]; exists {
			counts[ch]++
		} else {
			counts[ch] = 1
		}
	}

	i, j, total := 0, 0, len(t)
	var result string
	for j < len(s) {
		if count, exists := counts[s[j]]; exists {
			counts[s[j]]--
			if count > 0 {
				total--
			}
		}

		for i <= j && total == 0 {
			if result == "" || len(result) > j-i+1 {
				result = string(s[i : j+1])
			}

			if count, exists := counts[s[i]]; exists {
				counts[s[i]]++
				if count >= 0 {
					total++
				}
			}

			i++
		}

		j++
	}

	return result
}

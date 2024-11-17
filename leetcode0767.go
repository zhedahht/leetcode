/*
LeetCode 767: https://leetcode.com/problems/reorganize-string/description/
*/

package leetcode

func reorganizeString(s string) string {
	counts := make([]int, 26)
	most := 0
	mostCh := byte('a')
	for _, ch := range []byte(s) {
		counts[ch-'a']++
		if counts[ch-'a'] > most {
			most = counts[ch-'a']
			mostCh = ch
		}
	}

	others := len(s) - most
	if most > others+1 {
		return ""
	}

	result := make([]byte, len(s))
	i := 0
	for j := 0; j < most; j++ {
		result[i] = mostCh
		i += 2
	}

	for j, count := range counts {
		if byte(j)+byte('a') != mostCh {
			for k := 0; k < count; k++ {
				if i >= len(s) {
					i = 1
				}

				result[i] = byte(j + 'a')
				i += 2
			}
		}
	}

	return string(result)
}

/*
LeetCode 833: https://leetcode.com/problems/find-and-replace-in-string/
*/

package leetcode

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	min := func(m, n int) int {
		if m < n {
			return m
		}

		return n
	}

	m := make([]int, len(S))
	for i := range m {
		m[i] = -1
	}

	for i, v := range indexes {
		m[v] = i
	}

	result := make([]byte, 0)
	i := 0
	for i < len(S) {
		j := m[i]
		if j < 0 {
			result = append(result, S[i])
			i++
		} else {
			start, end := indexes[j], min(len(S), indexes[j]+len(sources[j]))
			sub := S[start:end]
			if sub == sources[j] {
				result = append(result, targets[j]...)
			} else {
				result = append(result, sub...)
			}

			i = end
		}
	}

	return string(result)
}

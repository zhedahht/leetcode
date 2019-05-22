/*
LeetCode 30: https://leetcode.com/problems/substring-with-concatenation-of-all-words/
*/

package leetcode

func findSubstring(s string, words []string) []int {
	wordToCount := make(map[string]int)
	total := 0
	for _, word := range words {
		total += len(word)
		if _, exists := wordToCount[word]; exists {
			wordToCount[word]++
		} else {
			wordToCount[word] = 1
		}
	}

	result := make([]int, 0)
	if len(s) < total || len(s) == 0 || total == 0 {
		return result
	}

	bytes, counts, length := []byte(s), len(words), len(words[0])
	for i := 0; i < length; i++ {
		clone := make(map[string]int)
		for key, val := range wordToCount {
			clone[key] = val
		}

		matched := 0
		for j := 0; j < counts-1; j++ {
			start := j*length + i
			word := string(bytes[start : start+length])
			if count, exists := clone[word]; exists {
				clone[word]--
				if count > 0 {
					matched++
				}
			}
		}

		for j := i + length*(counts-1); j <= len(s)-length; j += length {
			word := string(bytes[j : j+length])
			if count, exists := clone[word]; exists {
				clone[word]--
				if count > 0 {
					matched++
				}
			}

			start := j - length*(counts-1)
			if matched == counts {
				result = append(result, start)
			}

			word = string(bytes[start : start+length])
			if count, exists := clone[word]; exists {
				clone[word]++
				if count >= 0 {
					matched--
				}
			}
		}
	}

	return result
}

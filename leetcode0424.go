/*
LeetCode 424: https://leetcode.com/problems/longest-repeating-character-replacement/description/
*/

package leetcode

func characterReplacement(s string, k int) int {
	start, end := 0, -1
	counts := make(map[byte]int)
	result := 0
	for true {
		maxCount, total := 0, 0
		for _, count := range counts {
			maxCount = max(maxCount, count)
			total += count
		}

		if total-maxCount > k {
			counts[s[start]]--
			start++
			continue
		}

		result = max(result, end-start+1)
		end++
		if end >= len(s) {
			break
		}

		counts[s[end]]++
	}

	return result
}

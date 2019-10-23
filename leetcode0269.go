/*
LeetCode 269: https://leetcode.com/problems/alien-dictionary/
*/

package leetcode

import "fmt"

func alienOrder(words []string) string {
	charToPrev, charToNext := make(map[byte]map[byte]bool), make(map[byte]map[byte]bool)
	for _, word := range words {
		for _, ch := range []byte(word) {
			if _, exists := charToPrev[ch]; !exists {
				charToPrev[ch], charToNext[ch] = make(map[byte]bool), make(map[byte]bool)
			}
		}
	}

	for i := 0; i < len(words)-1; i++ {
		word1, word2 := words[i], words[i+1]
		for j := 0; j < len(word1) || j < len(word2); j++ {
			if j == len(word2) {
				return ""
			}

			if j == len(word1) {
				break
			}

			ch1, ch2 := byte(word1[j]), byte(word2[j])
			if ch1 != ch2 {
				charToNext[word1[j]][word2[j]], charToPrev[word2[j]][word1[j]] = true, true
				break
			}
		}
	}

	var queue []byte
	for ch, prevs := range charToPrev {
		if len(prevs) == 0 {
			delete(charToPrev, ch)
			queue = append(queue, ch)
		}
	}

	result := ""
	for len(queue) > 0 {
		ch := queue[0]
		result = fmt.Sprintf("%s%c", result, ch)
		queue = queue[1:]

		for next := range charToNext[ch] {
			prevsOfNext := charToPrev[next]
			delete(prevsOfNext, ch)
			if len(prevsOfNext) == 0 {
				delete(charToPrev, next)
				queue = append(queue, next)
			}
		}
	}

	if len(charToPrev) > 0 {
		return ""
	}

	return result
}

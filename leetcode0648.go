/*
LeetCode 648: https://leetcode.com/problems/replace-words/description/
*/

package leetcode

import (
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	words := strings.Split(sentence, " ")
	replaced := make([]string, 0)
	for _, word := range words {
		found := false
		for _, root := range dictionary {
			if strings.HasPrefix(word, root) {
				replaced = append(replaced, root)
				found = true
				break
			}
		}

		if !found {
			replaced = append(replaced, word)
		}
	}

	return strings.Join(replaced, " ")
}

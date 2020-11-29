/*
LeetCode 127: https://leetcode.com/problems/word-ladder/
*/

package leetcode

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	getNexts := func(word string, notVisited map[string]bool) []string {
		nexts := make([]string, 0)
		for i, c := range word {
			for ch := 'a'; ch <= 'z'; ch++ {
				if ch == c {
					continue
				}

				next := fmt.Sprintf("%s%c%s", word[:i], ch, word[i+1:])
				if notVisited[next] {
					nexts = append(nexts, next)
					delete(notVisited, next)
				}
			}
		}

		return nexts
	}

	queue1, queue2 := make([]string, 0), make([]string, 0)

	notVisited := make(map[string]bool)
	for _, word := range wordList {
		notVisited[word] = true
	}

	queue1 = append(queue1, beginWord)
	steps := 1
	for len(queue1) > 0 {
		word := queue1[0]
		queue1 = queue1[1:]
		if word == endWord {
			return steps
		}

		nexts := getNexts(word, notVisited)
		for _, next := range nexts {
			queue2 = append(queue2, next)
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]string, 0)
			steps++
		}
	}

	return 0
}

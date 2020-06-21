/*
LeetCode https://leetcode.com/problems/reorder-data-in-log-files/
*/

package leetcode

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	getFirstSpace := func(str string) int {
		for i, c := range str {
			if c == ' ' {
				return i
			}
		}

		return -1
	}

	isDigitLog := func(log string) bool {
		idx := getFirstSpace(log)
		return log[idx+1] >= '0' && log[idx+1] <= '9'
	}

	letterLogs, digitLogs := make([]string, 0), make([]string, 0)
	for _, log := range logs {
		if isDigitLog(log) {
			digitLogs = append(digitLogs, log)
		} else {
			letterLogs = append(letterLogs, log)
		}
	}

	sort.Slice(letterLogs, func(i, j int) bool {
		log1, log2 := letterLogs[i], letterLogs[j]
		idx1, idx2 := getFirstSpace(log1), getFirstSpace(log2)
		result := strings.Compare(log1[idx1+1:], log2[idx2+1:])
		if result != 0 {
			return result < 0
		}

		return strings.Compare(log1[:idx1], log2[:idx2]) < 0
	})

	return append(letterLogs, digitLogs...)
}

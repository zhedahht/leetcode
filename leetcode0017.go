/*
LeetCode 17: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
*/

package leetcode

func letterCombinations(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	queue := make(chan []string)
	count := len(digitMap[str[0]])
	for i := 0; i < count; i++ {
		go numToWordsCore(str, digitMap, 1, queue)
	}

	result := make([]string, 0)
	for i := 0; i < count; i++ {
		suffix := <-queue
		for _, word := range suffix {
			result = append(result, digitMap[str[0]][i:i+1]+word)
		}
	}

	return result
}

func numToWordsCore(str string, digitMap map[byte]string, index int, queue chan []string) {
	if index < len(str) {
		queueInner := make(chan []string)
		count := len(digitMap[str[index]])
		for i := 0; i < count; i++ {
			go numToWordsCore(str, digitMap, index+1, queueInner)
		}

		result := make([]string, 0)
		for i := 0; i < count; i++ {
			suffix := <-queueInner
			for _, word := range suffix {
				result = append(result, digitMap[str[index]][i:i+1]+word)
			}
		}

		queue <- result
	} else {
		queue <- []string{""}
	}
}

/*
LeetCode 472: https://leetcode.com/problems/concatenated-words/description/
*/

package leetcode

func findAllConcatenatedWordsInADict(words []string) []string {
	root := buildTrie(words)
	result := make([]string, 0)
	found := make(map[string]bool)
	notFound := make(map[string]bool)
	for _, word := range words {
		if isConcatenatedWord(word, root, root, found, notFound) {
			result = append(result, word)
		}
	}

	return result
}

func isConcatenatedWord(word string, root *trieNode, node *trieNode, found, notFound map[string]bool) bool {
	count := 0
	result := helper472(word, root, root, 0, found, notFound, &count)
	if result {
		found[word] = result
	}

	return result
}

func helper472(word string, root *trieNode, node *trieNode, index int, found, notFound map[string]bool, count *int) bool {
	if index == len(word) {
		if node.isWord {
			*count++
			if *count > 1 {
				return true
			}

			*count--
		}

		return false
	}

	if node.children[word[index]-'a'] == nil {
		return false
	}

	node = node.children[word[index]-'a']
	if helper472(word, root, node, index+1, found, notFound, count) {
		return true
	}

	if node.isWord {
		val, exists := found[word[index+1:]]
		if exists && val {
			return true
		}

		val, exists = notFound[word[index+1:]]
		if exists && val {
			return false
		}

		(*count)++
		val = helper472(word, root, root, index+1, found, notFound, count)
		if val {
			return val
		} else {
			notFound[word[index+1:]] = true
		}

		(*count)--
	}

	return false
}

/*
LeetCode 676: https://leetcode.com/problems/implement-magic-dictionary/description/
*/

package leetcode

type MagicDictionary struct {
	root *trieNode
}

// The function name should be Constructor. It's renamed to avoid conflict.
func Constructor676() MagicDictionary {
	return MagicDictionary{
		root: &trieNode{
			isWord:   false,
			children: [26]*trieNode{},
		},
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		node := this.root
		for _, ch := range word {
			index := int(ch - 'a')
			if node.children[index] == nil {
				node.children[index] = &trieNode{
					isWord:   false,
					children: [26]*trieNode{},
				}
			}

			node = node.children[index]
		}

		node.isWord = true
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return this.search([]rune(searchWord), this.root, 0)
}

func (this *MagicDictionary) search(word []rune, node *trieNode, change int) bool {
	if node == nil || change > 1 {
		return false
	}

	if len(word) == 0 {
		if node.isWord && change == 1 {
			return true
		} else {
			return false
		}
	}

	found := false
	for i := range node.children {
		newChange := change
		if int(word[0]-rune('a')) != i {
			newChange = change + 1
		}

		found = this.search(word[1:], node.children[i], newChange)
		if found {
			break
		}
	}

	return found
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

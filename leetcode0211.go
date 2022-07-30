/*
LeetCode 211: https://leetcode.com/problems/design-add-and-search-words-data-structure/
*/

package leetcode

func createTrieNode() *trieNode {
	return &trieNode{
		children: [26]*trieNode{},
	}
}

type WordDictionary struct {
	root *trieNode
}

// Constructor0211 should be Constructor. It's renamed to avoid conflicts.
func Constructor0211() WordDictionary {
	return WordDictionary{
		root: createTrieNode(),
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, ch := range word {
		if node.children[ch-'a'] == nil {
			node.children[ch-'a'] = createTrieNode()
		}
		node = node.children[ch-'a']
	}
	node.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.search(word, 0, this.root)
}

func (this *WordDictionary) search(word string, i int, n *trieNode) bool {
	if i == len(word) {
		return n.isWord
	}
	if word[i] == '.' {
		for ch := 'a'; ch <= 'z'; ch++ {
			if n.children[ch-'a'] == nil {
				continue
			}
			if this.search(word, i+1, n.children[ch-'a']) {
				return true
			}
		}
		return false
	}
	if n.children[word[i]-'a'] == nil {
		return false
	}
	return this.search(word, i+1, n.children[word[i]-'a'])
}

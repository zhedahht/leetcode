/*
LeetCode 208: https://leetcode.com/problems/implement-trie-prefix-tree/
*/

package leetcode

func newTrieNode() *trieNode {
	children := [26]*trieNode{}
	return &trieNode{
		children: children,
	}
}

type Trie struct {
	root *trieNode
}

// Constructor0208 should be Constructor. It's renamed to avoid conflicts.
func Constructor0208() Trie {
	return Trie{
		root: newTrieNode(),
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for i, r := range word {
		index := byte(r) - byte('a')
		if node.children[index] == nil {
			node.children[index] = newTrieNode()
		}
		node = node.children[index]
		if i == len(word)-1 {
			node.isWord = true
		}
	}
}

func (this *Trie) Search(word string) bool {
	node := this.root
	for _, r := range word {
		index := byte(r) - byte('a')
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, r := range prefix {
		index := byte(r) - byte('a')
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node != nil

}

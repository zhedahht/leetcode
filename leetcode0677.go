/*
LeetCode 677: https://leetcode.com/problems/map-sum-pairs/description/
*/

package leetcode

type trieNodeWithValue struct {
	value    int
	children [26]*trieNodeWithValue
}

type MapSum struct {
	root *trieNodeWithValue
}

// The function name should be Constructor. It's renamed to avoid conflicts.
func Constructor677() MapSum {
	return MapSum{
		root: &trieNodeWithValue{
			children: [26]*trieNodeWithValue{},
		},
	}
}

func (this *MapSum) Insert(key string, val int) {
	node := this.root
	for _, r := range key {
		index := int(r - 'a')
		if node.children[index] == nil {
			node.children[index] = &trieNodeWithValue{
				children: [26]*trieNodeWithValue{},
			}
		}

		node = node.children[index]
	}

	node.value = val
}

func (this *MapSum) Sum(prefix string) int {
	node := this.root
	for _, r := range prefix {
		index := int(r - 'a')
		if node.children[index] == nil {
			return 0
		}

		node = node.children[index]
	}

	return this.sum(node)
}

func (this *MapSum) sum(node *trieNodeWithValue) int {
	if node == nil {
		return 0
	}

	value := node.value
	for _, child := range node.children {
		value += this.sum(child)
	}

	return value
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

/*
LeetCode 173: https://leetcode.com/problems/binary-search-tree-iterator/
*/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	node  *TreeNode
	stack []*TreeNode
}

// NOTE: The function name should be Constructor. Rename is to avoid conflicts.
func Constructor173(root *TreeNode) BSTIterator {
	return BSTIterator{
		node: root,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	for this.node != nil {
		this.stack = append(this.stack, this.node)
		this.node = this.node.Left
	}

	this.node = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	val := this.node.Val
	this.node = this.node.Right

	return val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0 || this.node != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

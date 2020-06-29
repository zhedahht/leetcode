/*
LeetCode 919: https://leetcode.com/problems/complete-binary-tree-inserter/
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
type CBTInserter struct {
	root  *TreeNode
	queue []*TreeNode
}

// Constructor919 should be Constructor.
// Rename to avoid name conflicts.
func Constructor919(root *TreeNode) CBTInserter {
	inserter := CBTInserter{
		queue: make([]*TreeNode, 0),
	}

	q := make([]*TreeNode, 0)
	if root != nil {
		q = append(q, root)
		inserter.Insert(root.Val)
	}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.Left != nil {
			q = append(q, node.Left)
			inserter.Insert(node.Left.Val)
		}

		if node.Right != nil {
			q = append(q, node.Right)
			inserter.Insert(node.Right.Val)
		}
	}

	return inserter
}

func (this *CBTInserter) Insert(v int) int {
	node := &TreeNode{Val: v}
	this.queue = append(this.queue, node)
	if len(this.queue) == 1 {
		this.root = node
		return -1
	}

	parent := this.queue[0]
	if parent.Left == nil {
		parent.Left = node
	} else {
		parent.Right = node
		this.queue = this.queue[1:]
	}

	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

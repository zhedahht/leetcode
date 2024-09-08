/*
LeetCode 450: https://leetcode.com/problems/delete-node-in-a-bst/description/
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	dummy := &TreeNode{
		Val:   -9999999,
		Right: root,
	}

	parent, isLeft := dummy, true
	for parent != nil {
		if parent.Left != nil && parent.Left.Val == key {
			break
		} else if parent.Right != nil && parent.Right.Val == key {
			isLeft = false
			break
		} else if parent.Val < key {
			parent = parent.Right
		} else {
			parent = parent.Left
		}
	}

	if parent == nil {
		return root
	}

	deleted := parent.Right
	if isLeft {
		deleted = parent.Left
	}

	leftChild, rightChild := deleted.Left, deleted.Right
	if rightChild == nil {
		rightChild = leftChild
	} else {
		least := rightChild
		for least.Left != nil {
			least = least.Left
		}

		least.Left = leftChild
	}

	if isLeft {
		parent.Left = rightChild
	} else {
		parent.Right = rightChild
	}

	return dummy.Right
}

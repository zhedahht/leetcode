/*
LeetCode 105: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	return helper105(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func helper105(preorder, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[preStart]}
	rootInorderIndex := inStart
	for ; rootInorderIndex <= inEnd; rootInorderIndex++ {
		if inorder[rootInorderIndex] == root.Val {
			break
		}
	}

	leftLen := rootInorderIndex - inStart
	rightLen := inEnd - inStart - leftLen
	root.Left = helper105(preorder, inorder,
		preStart+1, preStart+leftLen,
		inStart, rootInorderIndex-1)
	root.Right = helper105(preorder, inorder,
		preEnd-rightLen+1, preEnd,
		rootInorderIndex+1, inEnd)
	return root
}

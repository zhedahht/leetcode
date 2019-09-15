/*
LeetCode 106: https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
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
// Note: The function name should be buildTree. Rename it to avoid name conflict.
func buildTree106(inorder []int, postorder []int) *TreeNode {
	return helper106(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func helper106(inorder, postorder []int, inStart, inEnd, postStart, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	root := &TreeNode{Val: postorder[postEnd]}
	rootInorderIndex := inStart
	for ; rootInorderIndex <= inEnd; rootInorderIndex++ {
		if inorder[rootInorderIndex] == root.Val {
			break
		}
	}

	leftLen := rootInorderIndex - inStart
	root.Left = helper106(inorder, postorder,
		inStart, rootInorderIndex-1,
		postStart, postStart+leftLen-1)
	root.Right = helper106(inorder, postorder,
		rootInorderIndex+1, inEnd,
		postStart+leftLen, postEnd-1)
	return root
}

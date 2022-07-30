/*
LeetCode 538: https://leetcode.com/problems/convert-bst-to-greater-tree/
*/

package leetcode

func convertBST(root *TreeNode) *TreeNode {
	node := root
	stack := make([]*TreeNode, 0)
	sum := 0
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Right
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += node.Val
		node.Val = sum

		node = node.Left
	}
	return root
}

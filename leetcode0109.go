/*
LeetCode 109: https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
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
func sortedListToBST(head *ListNode) *TreeNode {
	getNodeCount := func(head *ListNode) int {
		count := 0
		for head != nil {
			count++
			head = head.Next
		}

		return count
	}

	count := getNodeCount(head)
	root, _ := helper109(head, 0, count-1)
	return root
}

func helper109(head *ListNode, start, end int) (*TreeNode, *ListNode) {
	if start > end {
		return nil, head
	}

	if start == end {
		return &TreeNode{Val: head.Val}, head.Next
	}

	mid := (start + end) / 2
	left, head := helper109(head, start, mid-1)

	root := &TreeNode{Val: head.Val}
	root.Left = left

	head = head.Next
	right, head := helper109(head, mid+1, end)
	root.Right = right

	return root, head
}

/*
LeetCode 21: https://leetcode.com/problems/merge-two-sorted-lists/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			node, l1 = node.Next, l1.Next
		} else {
			node.Next = l2
			node, l2 = node.Next, l2.Next
		}
	}

	if l1 == nil {
		node.Next = l2
	} else {
		node.Next = l1
	}

	return dummy.Next
}

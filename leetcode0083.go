/*
LeetCode 83: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// NOTE: The name should be deleteDuplicates. Rename it to avoid conflicts with others.
func deleteDuplicates83(head *ListNode) *ListNode {
	node := head
	for node != nil {
		next := node.Next
		for next != nil && next.Val == node.Val {
			next = next.Next
		}

		node.Next = next
		node = next
	}

	return head
}

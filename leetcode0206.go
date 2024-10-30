/*
LeetCode 206: https://leetcode.com/problems/reverse-linked-list/description/
*/

package leetcode

func reverseList(head *ListNode) *ListNode {
	var reversedHead *ListNode
	var prev *ListNode
	node := head
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node

		reversedHead = node
		node = next
	}

	return reversedHead
}

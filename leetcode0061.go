/*
LeetCode 61: https://leetcode.com/problems/rotate-list/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	length := 0
	node := head
	for node != nil {
		length, node = length+1, node.Next
	}

	k = k % length
	if k == 0 {
		return head
	}

	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}

	newHead := slow.Next
	slow.Next, fast.Next = nil, head
	return newHead
}

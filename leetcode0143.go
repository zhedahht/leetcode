/*
LeetCode 143: https://leetcode.com/problems/reorder-list/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	head1, head2 := head, slow.Next
	slow.Next = nil

	var prev *ListNode
	node := head2
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	head2 = prev
	dummy := &ListNode{}
	node = dummy
	i := 0
	for head1 != nil || head2 != nil {
		if i%2 == 0 {
			node.Next = head1
			node, head1 = node.Next, head1.Next
		} else {
			node.Next = head2
			node, head2 = node.Next, head2.Next
		}

		i++
	}
}

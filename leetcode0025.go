/*
LeetCode 25: https://leetcode.com/problems/reverse-nodes-in-k-group/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	node := dummy
	for node != nil && node.Next != nil {
		first := node.Next
		i, cur := 1, first
		for i < k && cur != nil {
			cur = cur.Next
			i++
		}

		if cur == nil {
			break
		}

		last, after := cur, cur.Next
		cur = first
		var prev, next *ListNode
		for cur != after {
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		first.Next, node.Next = after, last
		node = first
	}

	return dummy.Next
}

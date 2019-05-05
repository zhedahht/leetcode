/*
LeetCode 24: https://leetcode.com/problems/swap-nodes-in-pairs/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	node := dummy
	for node != nil && node.Next != nil {
		n1 := node.Next
		var n2, n3 *ListNode
		if n1 != nil {
			n2 = n1.Next
		}

		if n2 != nil {
			n3 = n2.Next
			node.Next = n2
			n2.Next = n1
			n1.Next = n3
		}

		node = n1
	}

	return dummy.Next
}

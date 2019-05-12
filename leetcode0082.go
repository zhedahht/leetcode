/*
LeetCode 82: https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
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
func deleteDuplicates82(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	node := dummy
	for node != nil {
		next1, next2 := node.Next, node.Next
		for next2 != nil && next2.Val == next1.Val {
			next2 = next2.Next
		}

		if next1 != nil && next2 != next1.Next {
			node.Next = next2
		} else {
			node = next1
		}
	}

	return dummy.Next
}

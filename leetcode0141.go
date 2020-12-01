/*
LeetCode 141: https://leetcode.com/problems/linked-list-cycle/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head.Next
	var fast *ListNode
	if slow != nil {
		fast = slow.Next
	}

	for fast != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	return false
}

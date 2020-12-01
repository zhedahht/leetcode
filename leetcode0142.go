/*
LeetCode 142: https://leetcode.com/problems/linked-list-cycle-ii/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head.Next
	var fast *ListNode
	if slow != nil {
		fast = slow.Next
	}

	for fast != nil {
		if slow == fast {
			break
		}

		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	if fast == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return fast
}

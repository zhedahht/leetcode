/*
LeetCode 876: https://leetcode.com/problems/middle-of-the-linked-list/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, dummy
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow.Next
}

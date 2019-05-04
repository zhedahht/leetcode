/*
LeetCode 92: https://leetcode.com/problems/reverse-linked-list-ii/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	i := 0
	node := dummy
	for i < m-1 {
		node = node.Next
		i++
	}

	before := node
	node = node.Next
	first := node
	i++

	var prev, after, revHead *ListNode
	for i <= n {
		after = node.Next
		node.Next = prev
		prev = node
		if i == n {
			revHead = node
		}

		node = after
		i++
	}

	before.Next, first.Next = revHead, after
	return dummy.Next
}

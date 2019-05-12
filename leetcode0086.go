/*
LeetCode 86: https://leetcode.com/problems/partition-list/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	node1, node2 := dummy1, dummy2
	node := head
	for node != nil {
		if node.Val < x {
			node1.Next = node
			node1 = node1.Next
		} else {
			node2.Next = node
			node2 = node2.Next
		}

		node = node.Next
	}

	node1.Next, node2.Next = dummy2.Next, nil
	return dummy1.Next
}

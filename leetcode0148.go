/*
LeetCode 148: https://leetcode.com/problems/sort-list/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	end := head
	for end.Next != nil {
		end = end.Next
	}

	return helper148(head, end)
}

func helper148(head, end *ListNode) *ListNode {
	if head == end {
		return head
	}

	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != end {
		slow = slow.Next
		fast = fast.Next
		if fast != end {
			fast = fast.Next
		}
	}

	list1, list2 := head, slow.Next
	slow.Next, end.Next = nil, nil

	list1 = helper148(list1, slow)
	list2 = helper148(list2, end)

	node := dummy
	for list1 != nil || list2 != nil {
		if list1 == nil || (list2 != nil && list1.Val > list2.Val) {
			node.Next = list2
			list2, node = list2.Next, node.Next
		} else {
			node.Next = list1
			list1, node = list1.Next, node.Next
		}
	}

	return dummy.Next
}

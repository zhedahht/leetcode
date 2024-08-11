/*
LeetCode 445: https://leetcode.com/problems/add-two-numbers-ii/description/
*/

package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)

	carry := 0
	var prev *ListNode
	var first *ListNode
	for l1 != nil || l2 != nil {
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}

		node := &ListNode{
			Val: sum,
		}

		if prev != nil {
			prev.Next = node
		}

		if prev == nil {
			first = node
		}

		prev = node
	}

	if carry == 1 {
		prev.Next = &ListNode{
			Val: 1,
		}
	}

	return reverse(first)
}

func reverse(l *ListNode) *ListNode {
	node, last := l, l
	var prev *ListNode
	for node != nil {
		last = node
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	return last
}

/*
LeetCode 729: https://leetcode.com/problems/my-calendar-i/description/
*/

package leetcode

type calendarNode struct {
	start int
	end   int
	left  *calendarNode
	right *calendarNode
}

type MyCalendar struct {
	head *calendarNode
}

// The name should be Constructor. It's renamed to avoid conflict.
func Constructor729() MyCalendar {
	return MyCalendar{
		head: nil,
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	return this.insert(&this.head, start, end)
}

func (this *MyCalendar) insert(node **calendarNode, start int, end int) bool {
	if *node == nil {
		*node = &calendarNode{
			start: start,
			end:   end,
		}

		return true
	}

	if (*node).start >= end {
		return this.insert(&((*node).left), start, end)
	}

	if (*node).end <= start {
		return this.insert(&((*node).right), start, end)
	}

	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

/*
LeetCode 732: https://leetcode.com/problems/my-calendar-iii/description/
*/

package leetcode

import "fmt"

type calendarNode732 struct {
	startTime int
	endTime   int
	count     int
	left      *calendarNode732
	right     *calendarNode732
}

type MyCalendarThree struct {
	maxCount int
	root     *calendarNode732
}

// The name should be Constructor. It's renamed to avoid conflicts.
func Constructor732() MyCalendarThree {
	return MyCalendarThree{}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	this.insertCalendarNode(&this.root, startTime, endTime, 1)
	return this.maxCount
}

func (this *MyCalendarThree) insertCalendarNode(node **calendarNode732, startTime int, endTime int, count int) {
	if *node == nil {
		*node = &calendarNode732{
			startTime: startTime,
			endTime:   endTime,
			count:     count,
		}

		this.maxCount = max(this.maxCount, count)
		fmt.Println("new node", startTime, endTime, count)
		return
	}

	nodeStart, nodeEnd := (*node).startTime, (*node).endTime
	if nodeEnd <= startTime {
		this.insertCalendarNode(&((*node).right), startTime, endTime, count)
		return
	}

	if nodeStart >= endTime {
		this.insertCalendarNode(&((*node).left), startTime, endTime, count)
		return
	}

	if nodeStart != startTime {
		if nodeStart < startTime {
			this.insertCalendarNode(&((*node).left), nodeStart, startTime, (*node).count)
		} else {
			this.insertCalendarNode(&((*node).left), startTime, nodeStart, count)
		}
	}

	if nodeEnd != endTime {
		if nodeEnd > endTime {
			this.insertCalendarNode(&((*node).right), endTime, nodeEnd, (*node).count)
		} else {
			this.insertCalendarNode(&((*node).right), nodeEnd, endTime, count)
		}
	}

	(*node).startTime = max(nodeStart, startTime)
	(*node).endTime = min(nodeEnd, endTime)
	(*node).count += count
	this.maxCount = max(this.maxCount, (*node).count)
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */

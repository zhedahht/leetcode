/*
LeetCode: https://leetcode.com/problems/time-based-key-value-store/
*/

package leetcode

type TimeValue struct {
	time  int
	value string
}

type TimeMap struct {
	times map[string][]TimeValue
}

// Constructor981 should be Constructor.
// Rename it to avoid name conflicts.
func Constructor981() TimeMap {
	return TimeMap{
		times: make(map[string][]TimeValue),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	val := TimeValue{
		time:  timestamp,
		value: value,
	}

	this.times[key] = append(this.times[key], val)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	t := this.times[key]
	left, right := 0, len(t)-1
	for left <= right {
		mid := (left + right) / 2
		timeValue := t[mid]
		if timeValue.time <= timestamp {
			if mid == len(t)-1 || t[mid+1].time > timestamp {
				return t[mid].value
			}

			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ""
}

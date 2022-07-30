/*
LeetCode 692: https://leetcode.com/problems/top-k-frequent-words/
*/

package leetcode

import (
	"container/heap"
)

type wordFreq struct {
	word string
	freq int
}

func (f *wordFreq) Less(o *wordFreq) bool {
	if f.freq != o.freq {
		return f.freq < o.freq
	}
	return f.word > o.word
}

type freqQueue []*wordFreq

func (q *freqQueue) Len() int {
	return len(*q)
}

func (q *freqQueue) Less(i, j int) bool {
	itemI, itemJ := (*q)[i], (*q)[j]
	return itemI.Less(itemJ)
}

func (q *freqQueue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *freqQueue) Push(freq interface{}) {
	*q = append(*q, freq.(*wordFreq))
}

func (q *freqQueue) Pop() interface{} {
	freq := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return freq
}

func (q *freqQueue) Top() interface{} {
	return (*q)[0]
}

func topKFrequent(words []string, k int) []string {
	freqs := make(map[string]int)
	for _, word := range words {
		freqs[word]++
	}

	queue := make(freqQueue, 0)
	heap.Init(&queue)

	for w, f := range freqs {
		freq := &wordFreq{
			word: w,
			freq: f,
		}

		if queue.Len() < k {
			heap.Push(&queue, freq)
		} else {
			top := queue.Top().(*wordFreq)
			if top.Less(freq) {
				heap.Pop(&queue)
				heap.Push(&queue, freq)
			}
		}
	}

	result := make([]string, 0)
	for len(queue) > 0 {
		freq := heap.Pop(&queue)
		result = append(result, freq.(*wordFreq).word)
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

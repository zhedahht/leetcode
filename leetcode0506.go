/*
LeetCode 0506: https://leetcode.com/problems/relative-ranks/
*/

package leetcode

import (
	"fmt"
	"sort"
)

type NumToIndex struct {
	num   int
	index int
}

type NumToIndexSorter struct {
	numToIndices []NumToIndex
}

func (n *NumToIndexSorter) Len() int {
	return len(n.numToIndices)
}

func (n *NumToIndexSorter) Swap(i, j int) {
	n.numToIndices[i], n.numToIndices[j] = n.numToIndices[j], n.numToIndices[i]
}

func (n *NumToIndexSorter) Less(i, j int) bool {
	return n.numToIndices[i].num > n.numToIndices[j].num
}

func findRelativeRanks(nums []int) []string {
	numToIndices := make([]NumToIndex, len(nums))
	for index, num := range nums {
		numToIndices[index] = NumToIndex{num, index}
	}

	sort.Sort(&NumToIndexSorter{numToIndices: numToIndices})

	result := make([]string, len(nums))
	for i, numToIndex := range numToIndices {
		switch i {
		case 0:
			result[numToIndex.index] = "Gold Medal"
		case 1:
			result[numToIndex.index] = "Silver Medal"
		case 2:
			result[numToIndex.index] = "Bronze Medal"
		default:
			result[numToIndex.index] = fmt.Sprintf("%d", i+1)
		}
	}

	return result
}

/*
LeetCode 244: https://leetcode.com/problems/shortest-word-distance-ii/
*/

package leetcode

import "math"

type WordDistance struct {
	wordToIndices map[string][]int
}

// NOTE: it should be Constructor.
// Rename the function to avoid name conflicts with other.
func Constructor244(words []string) WordDistance {
	wordToIndices := make(map[string][]int)
	for i, word := range words {
		if indices, exist := wordToIndices[word]; exist {
			wordToIndices[word] = append(indices, i)
		} else {
			indices = make([]int, 0)
			wordToIndices[word] = append(indices, i)
		}
	}

	return WordDistance{
		wordToIndices: wordToIndices,
	}
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	indices1 := this.wordToIndices[word1]
	indices2 := this.wordToIndices[word2]
	minDist := 0x7fffffff
	for i, j := 0, 0; true; {
		dist := int(math.Abs(float64(indices1[i]) - float64(indices2[j])))
		minDist = int(math.Min(float64(minDist), float64(dist)))

		if i == len(indices1)-1 && j == len(indices2)-1 {
			break
		}

		if i == len(indices1)-1 || (j < len(indices2)-1 && indices1[i] > indices2[j]) {
			j++
		} else {
			i++
		}
	}

	return minDist
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Shortest(word1,word2);
 */

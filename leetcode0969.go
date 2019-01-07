/*
LeetCode 969: https://leetcode.com/problems/pancake-sorting/
*/

package leetcode

func pancakeSort(A []int) []int {
    result := make([]int, 0)
    for i := len(A) - 1; i >= 0; i-- {
        j := findLargest(A, i + 1)

        if j != i {
            flipFirst(A, j + 1)
            flipFirst(A, i + 1)
            
            result = append(result, j + 1)
            result = append(result, i + 1)
        }
    }
    
    return result
}

func findLargest(A []int, k int) int {
    largest := 0
    for i := 1; i < k; i++ {
        if A[i] > A[largest] {
            largest = i
        }
    }
    
    return largest
}

func flipFirst(A []int, k int) {
    i, j := 0, k - 1
    for i < j {
        A[i], A[j] = A[j], A[i]
        i, j = i + 1, j - 1
    }
}
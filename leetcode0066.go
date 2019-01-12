/*
LeetCode 66: https://leetcode.com/problems/plus-one/
*/

package leetcode

func plusOne(digits []int) []int {
    result := make([]int, 0)
    val := 1
    for i := len(digits) - 1; i >= 0; i-- {
        sum := val + digits[i]
        if sum >= 10 {
            sum -= 10
            val = 1
        } else {
            val = 0
        }
        
        result = append(result, sum)
    }
    
    if val > 0 {
        result = append(result, val)
    }
    
    for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
        result[i], result[j] = result[j], result[i]
    }
    
    return result
}
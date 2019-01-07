/*
LeetCode 263: https://leetcode.com/problems/ugly-number/
*/

package leetcode

func isUgly(num int) bool {
    if num < 1 {
        return false
    }
    
    for num > 1 {
        if num % 2 == 0 {
            num = num / 2
        } else if num % 3 == 0 {
            num = num / 3
        } else if num % 5 == 0 {
            num = num / 5
        } else {
            return false
        }
    }
    
    return true
}
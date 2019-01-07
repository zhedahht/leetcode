/*
LeetCode 970: https://leetcode.com/problems/powerful-integers/
*/

package leetcode

func powerfulIntegers(x int, y int, bound int) []int {
    set := make(map[int]bool)
    powX := 1
    for powX < bound {
        powY := 1
        for powY < bound {
            if (powX + powY <= bound) {
                set[powX + powY] = true
            }
            
            if y == 1 {
                break
            }
            
            powY *= y
        }
        
        if x == 1 {
            break
        }

        powX *= x
    }
    
    result := make([]int, 0)
    for k := range set {
        result = append(result, k)
    }
    
    return result
}
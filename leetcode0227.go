/*
LeetCode 227: https://leetcode.com/problems/basic-calculator-ii/description/
*/

package leetcode

func calculate(s string) int {
	nums := make([]int, 0)
	ops := make([]byte, 0)
	index := 0
	nums = append(nums, getNum(s, &index))
	for index < len(s) {
		op := getOp(s, &index)
		num := getNum(s, &index)

		if op == '+' || op == '-' {
			ops = append(ops, op)
			nums = append(nums, num)
		} else if op == '*' {
			nums[len(nums)-1] = nums[len(nums)-1] * num
		} else if op == '/' {
			nums[len(nums)-1] = nums[len(nums)-1] / num
		}
	}

	num := nums[0]
	for i, op := range ops {
		if op == '+' {
			num = num + nums[i+1]
		} else if op == '-' {
			num = num - nums[i+1]
		}
	}

	return num
}

func getOp(s string, index *int) byte {
	for s[*index] == ' ' {
		*index++
	}

	op := s[*index]
	*index++

	for s[*index] == ' ' {
		*index++
	}

	return op
}

func getNum(s string, index *int) int {
	for s[*index] == ' ' {
		*index++
	}

	num := 0
	for *index < len(s) && s[*index] >= '0' && s[*index] <= '9' {
		num = num*10 + int(s[*index]-'0')
		*index++
	}

	for *index < len(s) && s[*index] == ' ' {
		*index++
	}

	return num
}

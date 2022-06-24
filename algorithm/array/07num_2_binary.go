package main

import "fmt"

/**
 * @author: xaohuihui
 * @Path: go_study/algorithm/array/07num_2_binary.go
 * @Description:
 * @datetime: 2022/6/22 0:28:58
 * software: GoLand
**/
/* zheng整数转二进制
负整数 ：先是将对应的正整数转换成二进制后，对二进制取反，然后对结果再加一
*/

// 正整数转二进制
func positiveNum2Binary(num int) string {
	stack := make([]int, 0, 0)

	for num > 0 {
		stack = append(stack, num%2)
		num = num / 2
	}
	res := ""
	for i := len(stack) - 1; i >= 0; i-- {
		res += fmt.Sprintf("%d", stack[i])
	}
	return res
}


func main() {
	fmt.Println(positiveNum2Binary(8))
}

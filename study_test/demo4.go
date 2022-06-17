package main

import "fmt"

/**
 * @author: xaohuihui
 * @Path: go_study/study_test/demo4.go
 * @Description:
 * @datetime: 2022/5/22 11:21:55
 * software: GoLand
**/

type Stack struct {
	Queue []byte
}

func (s *Stack) push(value byte) {
	s.Queue = append(s.Queue, value)
}

func (s *Stack) pop() {
	s.Queue = s.Queue[:len(s.Queue)-1]
}

func removeOuterParentheses(s string) string {
	n := len(s)
	ans := make([]byte, 0)
	stack := &Stack{}
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			if len(stack.Queue) != 0 {
				ans = append(ans, s[i])
			}
			stack.push(s[i])
		} else if s[i] == ')' {
			stack.pop()
			if len(stack.Queue) != 0 {
				ans = append(ans, s[i])
			}
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

func main() {
	s := "(()())(())(()(()))"
	res := removeOuterParentheses(s)
	fmt.Println(res)
}

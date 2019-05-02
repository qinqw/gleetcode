package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(longestValidParentheses("()(()"))
	fmt.Println(longestValidParentheses("())(())"))
	fmt.Println(longestValidParentheses("()()(())"))
	fmt.Println(longestValidParentheses("()"))
}

func longestValidParentheses(s string) int {
	res := 0
	if len(s) == 0 || len(s) == 1 {
		return res
	}

	stack := Stack{}
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			if stack.IsEmpty() {
				start = i + 1
				continue
			}
			stack.Pop()

			if stack.IsEmpty() {
				res = maxInt(res, i-start+1)
			} else {
				top, _ := stack.Top()
				res = maxInt(res, i-top.(int))
			}
		}
	}
	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Stack the stack structure
type Stack []interface{}

// Len get the length of stack
func (stack Stack) Len() int {
	return len(stack)
}

// IsEmpty check the stack is empty
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

// Cap get che capacity of the stack
func (stack Stack) Cap() int {
	return cap(stack)
}

//Push node to stack
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

//Top get the top of stack
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

// Pop node from stack
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

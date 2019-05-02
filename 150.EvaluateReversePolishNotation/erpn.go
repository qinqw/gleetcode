package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"+", "2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	var stackT Stack
	res := 0
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			v1, err := stackT.Pop()
			v2, err2 := stackT.Pop()
			if err != nil {
				fmt.Println(err)
				break
			} else if err2 != nil {
				fmt.Println(err2)
				break
			}
			tmp := v1.(int) + v2.(int)
			stackT.Push(tmp)
		case "-":
			v1, err := stackT.Pop()
			v2, err2 := stackT.Pop()
			if err != nil {
				fmt.Println(err)
				break
			} else if err2 != nil {
				fmt.Println(err2)
				break
			}
			tmp := v1.(int) - v2.(int)
			stackT.Push(tmp)
		case "*":
			v1, err := stackT.Pop()
			v2, err2 := stackT.Pop()
			if err != nil {
				fmt.Println(err)
				break
			} else if err2 != nil {
				fmt.Println(err2)
				break
			}
			tmp := v1.(int) * v2.(int)
			stackT.Push(tmp)
		case "/":
			v1, err := stackT.Pop()
			v2, err2 := stackT.Pop()
			if err != nil {
				fmt.Println(err)
				break
			} else if err2 != nil {
				fmt.Println(err2)
				break
			}
			tmp := v1.(int) / v2.(int)
			stackT.Push(tmp)
		default:
			v, err := strconv.Atoi(tokens[i])
			if err != nil {
				fmt.Println(err)
				break
			}
			stackT.Push(v)
		}

	}
	v, _ := stackT.Pop()
	res = v.(int)
	return res
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

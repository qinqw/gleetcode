package main

import (
	"errors"
)

// TreeNode ,Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//BSTIterator the iterator of binary search tree
type BSTIterator struct {
	s Stack
}

// Constructor is constructor
func Constructor(root *TreeNode) BSTIterator {
	res := BSTIterator{}
	res.pushLeft(root)
	return res
}

func (h *BSTIterator) pushLeft(root *TreeNode) {
	if root != nil {
		h.s.Push(root)
		cur := root
		for cur.Left != nil {
			h.s.Push(cur.Left)
			cur = cur.Left
		}
	}
}

// Next  @return the next smallest number
func (h *BSTIterator) Next() int {
	v, _ := h.s.Top()
	top := v.(*TreeNode)
	h.s.Pop()
	h.pushLeft(top.Right)
	return top.Val
}

//HasNext @return whether we have a next smallest number
func (h *BSTIterator) HasNext() bool {
	return h.s.IsEmpty()
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

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

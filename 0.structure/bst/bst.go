package bst

import "github.com/qinqw/gleetcode/0.structure/stack"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode ,Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BSTIterator the iterator of binary search tree
type BSTIterator struct {
	s stack.Stack
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
	return !h.s.IsEmpty()
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

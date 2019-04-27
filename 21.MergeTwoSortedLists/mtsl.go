package main

import (
	"fmt"
	"strconv"
)

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Sting the link
func (t *ListNode) Sting() string {
	res := ""
	for {
		res = fmt.Sprint(res, t.Val)
		if t.Next == nil {
			break
		} else {
			t = t.Next
		}
	}
	return res
}

//PrintLink the link
func (t *ListNode) PrintLink() {
	for {
		//fmt.Printf(" %d,%p", t.Val, t.Next)
		fmt.Printf(" %d", t.Val)
		//if (t.Next == nil) || t.Next == t {
		if t.Next == nil {
			break
		} else {
			t = t.Next
		}
	}
	fmt.Printf("\n")
}

func makeSimpleLink(num int) *ListNode {
	var root, p *ListNode
	data := strconv.Itoa(num)
	lenData := len(data)
	for i := 0; i < lenData; i++ {
		b := fmt.Sprintf("%c", data[lenData-1-i])
		val, _ := strconv.Atoi(b)
		tmpNode := ListNode{Val: val, Next: nil}
		if i == 0 {
			root = &tmpNode
			p = &tmpNode
		} else {
			p.Next = &tmpNode
			p = &tmpNode
		}
	}
	return root
}

func main() {

	l1 := makeSimpleLink(421)
	l2 := makeSimpleLink(431)
	l1.PrintLink()
	l2.PrintLink()
	l3 := mergeTwoLists(l1, l2)

	l3.PrintLink()
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var root, p *ListNode
	//var tmp ListNode
	i := 0
	for l1 != nil && l2 != nil {
		if i == 0 {
			if l1.Val <= l2.Val {
				tmp := *l1
				p = &tmp
				root = p
				p.Next = nil
				l1 = l1.Next
			} else {
				tmp := *l2
				p = &tmp
				root = p
				p.Next = nil
				l2 = l2.Next
			}
		} else {
			if l1.Val <= l2.Val {
				tmp := *l1
				p.Next = &tmp
				p = p.Next
				p.Next = nil
				l1 = l1.Next
			} else {
				tmp := *l2
				p.Next = &tmp
				p = p.Next
				p.Next = nil
				l2 = l2.Next
			}
		}
		i++
	}
	if l1 != nil {
		if p == nil {
			root = l1
		} else {
			p.Next = l1
		}

	}
	if l2 != nil {
		if p == nil {
			root = l2
		} else {
			p.Next = l2
		}
	}
	return root
}

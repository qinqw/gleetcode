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
		fmt.Printf("%d", t.Val)
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

	l1 := makeSimpleLink(10)
	l2 := makeSimpleLink(11)

	l1.PrintLink()
	l2.PrintLink()

	l3 := AddTwoNumbers(l1, l2)
	l3.PrintLink()

	l1 = makeSimpleLink(0)
	l2 = makeSimpleLink(0)
	l1.PrintLink()
	l2.PrintLink()
	l3 = AddTwoNumbers(l1, l2)
	l3.PrintLink()

	l1 = makeSimpleLink(1)
	l2 = makeSimpleLink(1)
	l1.PrintLink()
	l2.PrintLink()
	l3 = AddTwoNumbers(l1, l2)
	l3.PrintLink()

	l1 = makeSimpleLink(122330)
	l2 = makeSimpleLink(1)
	l1.PrintLink()
	l2.PrintLink()
	l3 = AddTwoNumbers(l1, l2)
	l3.PrintLink()

	l1 = makeSimpleLink(5)
	l2 = makeSimpleLink(5)
	l1.PrintLink()
	l2.PrintLink()
	l3 = AddTwoNumbers(l1, l2)
	l3.PrintLink()
	// l1.PrintLink()
	// l2.PrintLink()

}

//AddTwoNumbers , Add two nubmbers
func AddTwoNumbers(tl1, tl2 *ListNode) *ListNode {
	var root, p *ListNode
	l1 := tl1
	l2 := tl2
	carryBit := 0
	i := 0

	for {
		tmpNode := ListNode{Val: 0, Next: nil}
		if i == 0 {
			root = &tmpNode
			p = &tmpNode
		}
		if l1 != nil && l2 != nil {
			sum := l1.Val + l2.Val + carryBit
			if sum < 10 {
				tmpNode.Val = sum
				carryBit = 0
			} else {
				tmpNode.Val = sum - 10
				carryBit = 1
			}

			p.Next = &tmpNode
			p = &tmpNode

			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			sum := l1.Val + carryBit
			if sum < 10 {
				tmpNode.Val = sum
				carryBit = 0
			} else {
				tmpNode.Val = sum - 10
				carryBit = 1
			}
			p.Next = &tmpNode
			p = &tmpNode
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			sum := l2.Val + carryBit
			if sum < 10 {
				tmpNode.Val = sum
				carryBit = 0
			} else {
				tmpNode.Val = sum - 10
				carryBit = 1
			}
			p.Next = &tmpNode
			p = &tmpNode
			l2 = l2.Next
		} else if l1 == nil && l2 == nil {
			if carryBit > 0 {
				tmpNode.Val = 1
				p.Next = &tmpNode
				p = &tmpNode
			}
			p.Next = nil
			break
		}
		i++
	}
	return root
}

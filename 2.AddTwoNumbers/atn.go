package main

import (
	"fmt"
	"strconv"
)

//LinkNode Definition for singly-linked list.
type LinkNode struct {
	Val  int
	Next *LinkNode
}

//Sting the link
func (t *LinkNode) Sting() string {
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
func (t *LinkNode) PrintLink() {
	for {
		fmt.Printf("%d", t.Val)
		if t.Next == nil {
			break
		} else {
			t = t.Next
		}
	}
	fmt.Printf("\n")
}

func main() {

	l1 := makeSimpleLink(68712)
	l2 := makeSimpleLink(789)

	l1.PrintLink()
	l2.PrintLink()

	l3 := AddTwoNumbers(l1, l2)
	l3.PrintLink()

	l1.PrintLink()
	l2.PrintLink()

}

func makeSimpleLink(num int) *LinkNode {
	var root, p *LinkNode
	data := strconv.Itoa(num)
	for i := len(data) - 1; i >= 0; i-- {
		b := fmt.Sprintf("%c", data[i])
		val, _ := strconv.Atoi(b)
		tmpNode := LinkNode{Val: val, Next: nil}
		if i == (len(data) - 1) {
			root = &tmpNode
			p = &tmpNode
		} else {
			p.Next = &tmpNode
			p = &tmpNode
		}
	}
	return root
}

//AddTwoNumbers , Add two nubmbers
func AddTwoNumbers(tl1, tl2 *LinkNode) *LinkNode {
	var root, p *LinkNode
	l1 := tl1
	l2 := tl2
	carryBit := 0
	i := 0
	for {
		tmpNode := LinkNode{Val: 0, Next: nil}
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
			break
		}
		i++
	}
	return root
}

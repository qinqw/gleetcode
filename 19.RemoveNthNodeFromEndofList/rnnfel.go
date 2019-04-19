package main

import (
	"fmt"
)

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	lk := makeSinglyLinked([]int{1, 2, 3, 4, 5})
	printSinglyLinked(lk)
	lk = removeNthFromEnd(lk, 3)
	printSinglyLinked(lk)
}

func makeSinglyLinked(nums []int) *ListNode {
	var root, p *ListNode
	for i, v := range nums {
		node := ListNode{
			Val: v,
		}
		if i == 0 {
			root = &node
			p = &node
		} else {
			p.Next = &node
			p = p.Next
		}
	}
	return root
}

func printSinglyLinked(l *ListNode) {
	p := l
	for p != nil {
		if p.Next == nil {
			fmt.Printf("%d", p.Val)
		} else {
			fmt.Printf("%d->", p.Val)
		}

		p = p.Next
	}
	fmt.Println("")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p *ListNode
	p = head
	list := []*ListNode{}
	i := 0
	for head != nil {
		i++
		list = append(list, head)
		head = head.Next
	}

	if i == n {
		p = list[0].Next
	} else {
		list[i-n-1].Next = list[i-n].Next
	}

	return p
}

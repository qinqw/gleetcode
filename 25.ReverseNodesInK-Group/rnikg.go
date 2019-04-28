package main

import "fmt"

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
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

func main() {
	k := 3
	lk := makeSinglyLinked([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	l := reverseKGroup(lk, k)
	printSinglyLinked(l)

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	tmp := []*ListNode{}
	var root, p1 *ListNode
	if head == nil {
		return head
	}
	p1 = head
	l := 0
	for p1 != nil {
		tmp = append(tmp, p1)
		p1 = p1.Next
		l++
	}
	if l < k {
		return head
	}

	//fmt.Printf("len l,k:%d,%d\n", l, k)
	for i := 0; i < l/k; i++ {
		for j := 0; j < k/2; j++ {
			//fmt.Printf("swap x,y:%d,%d\n", k*i+j, k*i+k-1-j)
			tmp[k*i+j], tmp[k*i+k-1-j] = tmp[k*i+k-1-j], tmp[k*i+j]
		}
	}
	for j := 0; j < len(tmp); j++ {
		if j < len(tmp)-1 {
			tmp[j].Next = tmp[j+1]
		} else {
			tmp[j].Next = nil
		}
	}
	root = tmp[0]

	return root
}

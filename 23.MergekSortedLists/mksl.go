package main

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var root, p *ListNode
	for i := 0; i < len(lists); i++ {
		p = mergeTwoLists(p, lists[i])
	}
	root = p
	return root
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

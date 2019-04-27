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
func swapPairs(head *ListNode) *ListNode {
	tmp := []*ListNode{}
	var root, p1 *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	p1 = head
	for p1 != nil {
		tmp = append(tmp, p1)
		p1 = p1.Next
	}

	for i := 0; i < len(tmp)/2; i++ {
		tmp[2*i], tmp[2*i+1] = tmp[2*i+1], tmp[2*i]
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

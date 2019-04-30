package main

func main() {

}

// MinStack the stack can return the min node
type MinStack struct {
	s []int
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	var res MinStack
	return res
}

func (h *MinStack) Push(x int) {
	h.s = append(h.s, x)
}

func (h *MinStack) Pop() int {
	theStack := h.s
	// if len(theStack) == 0 {
	// 	return nil
	// }
	value := theStack[len(theStack)-1]
	h.s = theStack[:len(theStack)-1]
	return value
}

func (h *MinStack) Top() int {
	// if len(h.s) == 0 {
	// 	return nil
	// }
	return h.s[len(h.s)-1]
}

func (h *MinStack) GetMin() int {
	if len(h.s) == 1 {
		return h.s[0]
	}
	min := h.s[0]
	for i := 0; i < len(h.s); i++ {
		if min > h.s[i] {
			min = h.s[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

package main

//MyCircularQueue is circular queue
type MyCircularQueue struct {
	queue []int
	caps  int
	lens  int
}

// Constructor , Initialize your data structure here.
// Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		caps: k,
	}
}

// EnQueue , Insert an element into the circular queue.
// Return true if the operation is successful.
func (h *MyCircularQueue) EnQueue(value int) bool {
	res := false
	if h.lens < h.caps {
		h.lens++
		h.queue = append(h.queue, value)
		res = true
	}
	return res
}

//DeQueue , Delete an element from the circular queue.
//Return true if the operation is successful.
func (h *MyCircularQueue) DeQueue() bool {
	res := false
	if h.lens > 0 {
		h.queue = h.queue[1:]
		h.lens--
		res = true
	}
	return res
}

// Front , Get the front item from the queue.
func (h *MyCircularQueue) Front() int {
	return h.queue[0]
}

// Rear Get the last item from the queue.
func (h *MyCircularQueue) Rear() int {
	res := -1
	if h.lens > 0 {
		res = h.queue[h.lens-1]
	}
	return res
}

// IsEmpty , Checks whether the circular queue is empty or not. */
func (h *MyCircularQueue) IsEmpty() bool {
	return h.lens == 0
}

// IsFull , Checks whether the circular queue is full or not.
func (h *MyCircularQueue) IsFull() bool {
	return h.lens == h.caps
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

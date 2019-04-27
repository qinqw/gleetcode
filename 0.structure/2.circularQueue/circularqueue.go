package main

//MyCircularQueue is circular queue
type MyCircularQueue struct {
	queue []int
	count int
}

// Constructor , Initialize your data structure here.
// Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	return *NewMyCircularQueue(k)
}

// NewMyCircularQueue , Initialize your data structure here.
// Set the size of the queue to be k.
func NewMyCircularQueue(k int) *MyCircularQueue {
	// res := MyCircularQueue{
	// 	queue: []int{},
	// 	count: k,
	// }
	//res.count = 3
	return &MyCircularQueue{
		queue: []int{},
		count: k,
	}
}

// EnQueue , Insert an element into the circular queue.
// Return true if the operation is successful.
func (h *MyCircularQueue) EnQueue(value int) bool {
	res := false
	if h.count < len(h.queue) {
		h.queue = append(h.queue, value)
		res = true
	}
	return res
}

//DeQueue , Delete an element from the circular queue.
//Return true if the operation is successful.
func (h *MyCircularQueue) DeQueue() bool {
	res := false
	if h.count < len(h.queue) {
		h.queue = h.queue[1:]
		res = true
	}
	return res
}

// Front , Get the front item from the queue.
func (h *MyCircularQueue) Front() int {
	return 0
}

// Rear Get the last item from the queue.
func (h *MyCircularQueue) Rear() int {
	return len(h.queue) - 1
}

// IsEmpty , Checks whether the circular queue is empty or not. */
func (h *MyCircularQueue) IsEmpty() bool {
	// res := true {
	// 	if len(h.queue)
	// }
	return (len(h.queue) > 0)
}

// IsFull , Checks whether the circular queue is full or not.
func (h *MyCircularQueue) IsFull() bool {
	return len(h.queue) == h.count
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

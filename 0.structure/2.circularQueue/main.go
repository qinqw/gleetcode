package main

import (
	"fmt"
)

func main() {
	cq := Constructor(3)
	//cq := MyCircularQueue(3)
	//cq = MyCircularQueue.Constructor(3)
	fmt.Println(cq)
	cq.EnQueue(3)
	fmt.Println(cq.IsFull())
	fmt.Println(cq.IsEmpty())
	cq.EnQueue(2)
	cq.EnQueue(1)
	cq.EnQueue(2)
	fmt.Println(cq.IsFull())
	fmt.Println(cq.IsEmpty())
	fmt.Println(cq.Rear())
	cq.DeQueue()
	cq.DeQueue()
	cq.DeQueue()
	fmt.Println(cq.IsFull())
	fmt.Println(cq.IsEmpty())
	fmt.Println(cq)
}

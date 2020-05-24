package solutions

// https://leetcode-cn.com/problems/design-circular-deque/

/*
array implementation, key parts are as followed:
when apply capacity, use capacity+1 instead of capacity to avoid ambiguous between full and empty
when tail == head, is empty
when (tail + 1)%capacity == head, is full
when doing add front, head - 1, then do parse
when doing add rear, do parse first, then tail + 1
*/

// // MyCircularDeque used to represent 0641 struct, inner has a circular linked list
// type MyCircularDeque struct {
// 	items    []int
// 	head     int
// 	tail     int
// 	capacity int
// }

// // Constructor Initialize your data structure here. Set the size of the deque to be k.
// func Constructor(k int) MyCircularDeque {
// 	c := k + 1
// 	return MyCircularDeque{items: make([]int, c), capacity: c}
// }

// // InsertFront Adds an item at the front of Deque. Return true if the operation is successful.
// func (d *MyCircularDeque) InsertFront(value int) bool {
// 	if d.IsFull() {
// 		return false
// 	}

// 	d.head = d.getValidIndex(d.head - 1)
// 	d.items[d.head] = value
// 	return true
// }

// // InsertLast Adds an item at the rear of Deque. Return true if the operation is successful.
// func (d *MyCircularDeque) InsertLast(value int) bool {
// 	if d.IsFull() {
// 		return false
// 	}

// 	d.items[d.tail] = value
// 	d.tail = d.getValidIndex(d.tail + 1)

// 	return true
// }

// // DeleteFront Deletes an item from the front of Deque. Return true if the operation is successful.
// func (d *MyCircularDeque) DeleteFront() bool {
// 	if d.IsEmpty() {
// 		return false
// 	}

// 	d.head = d.getValidIndex(d.head + 1)
// 	return true
// }

// // DeleteLast Deletes an item from the rear of Deque. Return true if the operation is successful.
// func (d *MyCircularDeque) DeleteLast() bool {
// 	if d.IsEmpty() {
// 		return false
// 	}

// 	d.tail = d.getValidIndex(d.tail - 1)
// 	return true
// }

// // GetFront Get the front item from the deque.
// func (d *MyCircularDeque) GetFront() int {
// 	if d.IsEmpty() {
// 		return -1
// 	}
// 	return d.items[d.head]
// }

// // GetRear  Get the last item from the deque.
// func (d *MyCircularDeque) GetRear() int {
// 	if d.IsEmpty() {
// 		return -1
// 	}
// 	return d.items[d.getValidIndex(d.tail-1)]
// }

// // IsEmpty Checks whether the circular deque is empty or not.
// func (d *MyCircularDeque) IsEmpty() bool {
// 	return d.tail == d.head
// }

// // IsFull Checks whether the circular deque is full or not.
// func (d *MyCircularDeque) IsFull() bool {
// 	return d.getValidIndex(d.tail+1) == d.head
// }

// func (d *MyCircularDeque) getValidIndex(rawIndex int) int {
// 	return (rawIndex + d.capacity) % d.capacity
// }

/*
linkedlist implementation, key parts are as followed:
when adding new node, first take care the new node prev and next ,then take care the linked node include (prev and next)
when checking is full, use size == capacity
when checking is empty, use d.head.next == d.tail or d.tail.prev == d.head
*/

// MyCircularDeque used to represent 0641 struct, inner has a circular linked list
type MyCircularDeque struct {
	head     *circularNode
	tail     *circularNode
	size     int
	capacity int
}

type circularNode struct {
	prev  *circularNode
	next  *circularNode
	value int
}

// Constructor Initialize your data structure here. Set the size of the deque to be k.
func Constructor(k int) MyCircularDeque {
	head := &circularNode{value: -1}
	tail := &circularNode{value: -1}
	head.next = tail
	tail.prev = head
	return MyCircularDeque{head: head, tail: tail, capacity: k}
}

// InsertFront Adds an item at the front of Deque. Return true if the operation is successful.
func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}

	newNode := &circularNode{value: value}

	newNode.next = d.head.next
	newNode.prev = d.head

	newNode.next.prev = newNode
	newNode.prev.next = newNode
	d.size++
	return true
}

// InsertLast Adds an item at the rear of Deque. Return true if the operation is successful.
func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}

	newNode := &circularNode{value: value}

	newNode.next = d.tail
	newNode.prev = d.tail.prev

	newNode.prev.next = newNode
	newNode.next.prev = newNode
	d.size++
	return true
}

// DeleteFront Deletes an item from the front of Deque. Return true if the operation is successful.
func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}

	firstNode := d.head.next

	d.head.next = firstNode.next
	firstNode.next.prev = d.head

	firstNode.next, firstNode.prev = nil, nil
	d.size--

	return true
}

// DeleteLast Deletes an item from the rear of Deque. Return true if the operation is successful.
func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}

	lastNode := d.tail.prev

	d.tail.prev = lastNode.prev
	lastNode.prev.next = d.tail

	lastNode.next, lastNode.prev = nil, nil
	d.size--

	return true
}

// GetFront Get the front item from the deque.
func (d *MyCircularDeque) GetFront() int {
	return d.head.next.value
}

// GetRear  Get the last item from the deque.
func (d *MyCircularDeque) GetRear() int {
	return d.tail.prev.value
}

// IsEmpty Checks whether the circular deque is empty or not.
func (d *MyCircularDeque) IsEmpty() bool {
	return d.head.next == d.tail
}

// IsFull Checks whether the circular deque is full or not.
func (d *MyCircularDeque) IsFull() bool {
	return d.capacity == d.size
}

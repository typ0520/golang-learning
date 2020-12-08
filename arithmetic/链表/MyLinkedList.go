package main

type ListNode struct {
	val  int
	next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	tail *ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{head: nil, tail: nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.head == nil {
		return -1
	}
	tail, i := this.head, 0
	for tail != nil && i < index {
		tail = tail.next
		i++
	}
	if i == index && tail != nil {
		return tail.val
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &ListNode{val: val, next: this.head}
	if this.tail == nil {
		this.tail = this.head
	}
}

// /** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{val: val, next: nil}
	if this.head == nil {
		this.head = node
	} else {
		this.tail.next = node
	}
	this.tail = node
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if this.head == nil || index < 0 {
		return
	}
	table := make(map[int]*ListNode)
	tail, i := this.head, 0
	for tail != nil && i < index {
		table[i] = tail
		tail = tail.next
		i++
	}
	if index > i {
		return
	}
	if tail == nil {
		table[i-1].next = &ListNode{val: val, next: nil}
		this.tail = table[i-1].next
	} else {
		table[i-1].next = &ListNode{val: val, next: tail}
	}
}

// /** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.head == nil || index < 0 {
		return
	}
	table := make(map[int]*ListNode)
	tail, i := this.head, 0
	for tail != nil && i <= index {
		table[i] = tail
		tail = tail.next
		i++
	}
	removeElement := table[index]
	if removeElement == nil {
		return
	}
	prev := table[index-1]
	var next *ListNode
	if removeElement != nil {
		next = removeElement.next
	}
	if prev == nil && next == nil {
		this.head = nil
		this.tail = nil
	} else {
		if prev == nil {
			this.head = next
		} else {
			prev.next = next
		}
	}
}

func (this *MyLinkedList) ValueList() []int {
	if this.head == nil {
		return []int{}
	}
	var result []int
	tail := this.head
	for tail != nil {
		result = append(result, tail.val)
		tail = tail.next
	}
	return result
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

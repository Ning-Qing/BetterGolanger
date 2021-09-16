/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */

// @lc code=start
type MyLinkedList struct {
	size int
	head *Node
}

type Node struct {
	Val  int
	Next *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	// head 指向虚拟头节点
	return MyLinkedList{size: 0, head: &Node{}}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.size-1 {
		return - 1
	}
	cur := this.head
	// 这里是i<=index
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0,val)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size,val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	if index < 0 {
		index = 0
	}
	cur := this.head
	// 这里是i<index
	for i := 0; i < index; i++ {
		cur =cur.Next
	}
	node := &Node{Val:val,Next:cur.Next}
	cur.Next = node
	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > this.size-1 {
		return
	}
	cur := this.head
	for i := 0; i < index; i++ {
		cur =cur.Next
	}
	cur.Next=cur.Next.Next
	this.size--
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
// @lc code=end


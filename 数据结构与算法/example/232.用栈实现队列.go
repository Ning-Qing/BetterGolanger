/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	inStack,outStack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{inStack:make([]int,0,100),outStack:make([]int,0,100)}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.inStack = append(this.inStack,x)
}

func (this *MyQueue) inToOut(){
	for len(this.inStack)>0{
		this.outStack = append(this.outStack,this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}	
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	// 当outStack为空，将inStack数据依次出栈放入outStack
	if len(this.outStack)==0{
			this.inToOut()
	}
	// 此时outStack是inStack的逆序
	res := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.outStack)==0{
		this.inToOut()
	}
	return this.outStack[len(this.outStack)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.inStack)==0 && len(this.outStack)==0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end


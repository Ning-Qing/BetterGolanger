/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	queue1,queue2 []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{queue1:make([]int,0,100),queue2:make([]int,0,100)}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	// 此时queue1存了所有数，queue2为空
	// 后入 如果要先出 就需要将其他数堆在它后面
	this.queue2 = append(this.queue2,x)
	// 将其他数堆在后面，queue2存了所有数，queue1 为空
	for len(this.queue1)>0 {
		this.queue2 = append(this.queue2,this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	// 交换,恢复成初始状态
	this.queue2,this.queue1 =this.queue1,this.queue2
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := this.queue1[0]
	this.queue1=this.queue1[1:]
	return res
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue1[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue1)==0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end


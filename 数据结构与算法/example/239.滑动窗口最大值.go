/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	// 构建单调递减队列
	queue := make([]int,0,len(nums))
	push := func(value int){
		// 入队时如若要保持单调性
		// 应拿入队元素和队列末端元素比较
		// 如果大于，则将末端元素弹出
		// 直到入队元素小于末端元素
		for len(queue) != 0&&value>queue[len(queue)-1]{
			queue = queue[:len(queue)-1]
		}
		queue =append(queue,value)
	}
	pop := func(value int){
		if len(queue) != 0&&value == queue[0]{
			queue = queue[1:]
		}
	}

	for i:=0;i<k;i++{
		push(nums[i])
	}
	res := make([]int,0,len(nums))
	res=append(res,queue[0])
	for i:=k;i<len(nums);i++{
		pop(nums[i-k])
		push(nums[i])
		res = append(res,queue[0])
	}
	return res
}
// @lc code=end


/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	slow,fast,length:=0,0,len(nums)
	for ;fast<length;{
		if nums[fast] == val{
			fast++
			continue
		}
		// slow的左边因该是已移除元素的数组
		nums[slow]=nums[fast]
		slow++
		fast++
	}
	return slow
}
// @lc code=end


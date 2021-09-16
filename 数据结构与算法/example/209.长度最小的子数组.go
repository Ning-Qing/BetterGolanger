/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	start, end := 0, 0
	length, sum := 0, 0
	win := 50000
	for ; end < len(nums); end++ {
		sum += nums[end]
		// 满足子数组和大于target
		for sum >= target {
			// 计算子序列长度
			length = end - start + 1
			if length < win {
				win = length
			}
			// 窗口缩小
			sum -= nums[start]
			start++
		}
		// 窗口增大
	}
	if win == 50000 {
		win = 0
	}
	return win
}

// @lc code=end


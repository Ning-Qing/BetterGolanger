/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	// target一定存在在[left,right)
	for ;left<right;{
		middle := (left+right)/2
		if nums[middle]==target{
			return middle
		}else if nums[middle]<target{
			// 维持target一定存在在[left,right)
			left = middle+1
		}else{
			right = middle
		}
	}
	// 如果找不到target，返回right
	// target不在[left,right),那有序插入位置一定是right
	return right
}
// @lc code=end

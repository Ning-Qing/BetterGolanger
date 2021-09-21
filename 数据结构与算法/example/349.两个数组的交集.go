/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _,v := range nums1{
		m[v]=1
	}
	for _,v := range nums2{
		if _,ok:=m[v];ok{
			m[v]++
		}
	}
	res := make([]int,0)
	for k,v:=range m{
		if v>1{
			res=append(res,k)
		}
	}
	return res
}
// @lc code=end


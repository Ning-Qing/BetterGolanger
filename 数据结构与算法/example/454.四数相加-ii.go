/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count :=0
	map1:=make(map[int]int)
	for _,v1 := range nums1{
		for _,v2 := range nums2{
			value := v1+v2
			if _,ok:=map1[value];!ok{
				map1[value]=1
			}else{
				map1[value]+=1
			}
		}
	}

	for _,v1 := range nums3{
		for _,v2 := range nums4{
			value := v1+v2
			if _,ok:=map1[-value];ok{
				count += map1[-value]
			}
		}
	}

	return count
}
// @lc code=end


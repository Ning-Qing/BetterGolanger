/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func getSum(n int)int{
	res :=0
	for n >0 {
		res +=(n%10)*(n%10)
		n/=10
	}
	return res
}

func isHappy(n int) bool {
	m := make(map[int]int)
	for sum:=getSum(n);sum!=1;sum=getSum(sum){
		if _,ok := m[sum];ok{
			return false
		}else{
			m[sum]=0
		}
	}
	return true
}
// @lc code=end


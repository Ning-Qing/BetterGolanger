/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte)  {
	a,b:=0,len(s)-1
	for a<b {
		s[a],s[b]=s[b],s[a]
		a++
		b--
	}
}
// @lc code=end


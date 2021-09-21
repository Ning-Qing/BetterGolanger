/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 小写英文字符26个
	// 从0逐个映射 0->a
	arr := [26]int{}
	for i:=0;i<len(s);i++{
		// [s[i]-'a'] 借用ASCII编码映射
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}
	for _,v:=range arr{
		if v !=0{
			return false
		}
	}
	return true
}
// @lc code=end


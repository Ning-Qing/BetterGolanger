/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	m :=map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	b := []byte(s)
	stack := make([]byte,0,len(s))
	for len(b)>0{
		char := b[len(b)-1]
		if len(stack)==0 || m[stack[len(stack)-1]]!=char{
			stack=append(stack,char)
			b = b[:len(b)-1]
		}else{
			stack=stack[:len(stack)-1]
			b = b[:len(b)-1]
		}
	}
	if len(stack)==0&&len(b)==0{
		return true
	}
	return false
}
// @lc code=end


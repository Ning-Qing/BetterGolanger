/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote)>len(magazine){
		return false
	}

	// 使用数组以节约空间
	hashMap :=[26]int{0}

	for _,v :=range magazine{
		hashMap[v-'a']+=1
	}

	for _,v :=range ransomNote{
		if hashMap[v-'a']-=1;hashMap[v-'a']<0{
			return false
		}
	}

	return true
}
// @lc code=end


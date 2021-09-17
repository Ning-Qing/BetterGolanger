/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 1->2->nil
	// nil->1->2->nil
	// nil<-1<-2<-nil
	var left,right *ListNode
	left = nil
	right = head
	for right != nil {
		tmp := right.Next
		right.Next = left
		left = right
		right = tmp
	}
	return left
}

// @lc code=end


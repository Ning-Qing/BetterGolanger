/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// 创建一个头节点
	h := &ListNode{Val:0,Next:head}
	for node:=h;node.Next != nil;{
		if node.Next.Val == val{
			node.Next = node.Next.Next
		}else {
			node = node.Next
		}
	}
	return h.Next
}
// @lc code=end


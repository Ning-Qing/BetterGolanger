/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		// 快慢指针分别以2节点和1节点速度前进
		slow = slow.Next
		fast = fast.Next.Next
		// 如果两指针相遇，那么一定在环内相遇
		if slow == fast {
			// 两指针分别从头和相遇节点出发，相遇时就是环的入口节点
			index1, index2 := head, slow
			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}
			return index1
		}
	}
	return nil
}

// @lc code=end


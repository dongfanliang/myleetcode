package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表反转
func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode = nil

	for cur != nil {
		// tmpNext = cur.Next
		// cur.Next = prev
		// prev = cur
		// cur = temNext
		// golang的多重赋值特性，可以简化为：
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	return prev
}

// 是否有环
/*
* 有3解法
* 1. 硬解决，判断是否有nil
* 2. 使用map，循环链表，判断是否在map中
* 3. 快慢指针，判断是否会重合
 */
func hasCycle(head *ListNode) bool {
	fast, slow := head, head

	for fast != nil && slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

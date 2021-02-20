package main

import "fmt"

type ListNode struct {
    Val int
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

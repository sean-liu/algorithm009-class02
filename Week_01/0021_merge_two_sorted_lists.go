package solutions

// https://leetcode-cn.com/problems/merge-two-sorted-lists/

// 迭代法, 时间 O(min(len(l1), len(2)))
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	prev := head

	getNextAndAppend := func(next *ListNode) *ListNode {
		prev.Next = next
		prev = next
		return next.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			l2 = getNextAndAppend(l2)
		} else {
			l1 = getNextAndAppend(l1)
		}
	}

	if l1 != nil {
		prev.Next = l1
	} else if l2 != nil {
		prev.Next = l2
	}

	return head.Next
}

// 递归法，时间 O(m+n), 空间O(m+n), 不推荐

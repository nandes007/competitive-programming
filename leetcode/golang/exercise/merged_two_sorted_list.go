package exercise

import "github.com/nandes007/leetcode/golang/model"

func MergeTwoLists(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	dummyHead := &model.ListNode{}
	tail := dummyHead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}

		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return dummyHead.Next
}

func SortList(head *model.ListNode) *model.ListNode {
	dummyHead := &model.ListNode{}
	dummyHead.Next = head

	for current := head; current != nil && current.Next != nil; {
		if current.Val <= current.Next.Val {
			current = current.Next
			continue
		}

		next := current.Next
		current.Next = next.Next

		insertPosition := dummyHead
		for insertPosition.Next.Val <= next.Val && insertPosition.Next != current {
			insertPosition = insertPosition.Next
		}

		next.Next = insertPosition.Next
		insertPosition.Next = next
	}

	return dummyHead.Next
}

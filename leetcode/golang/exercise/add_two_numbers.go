package exercise

import "github.com/nandes007/leetcode/golang/model"

func AddTwoNumbers(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	dummy := &model.ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &model.ListNode{Val: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &model.ListNode{Val: carry}
	}

	return dummy.Next
}

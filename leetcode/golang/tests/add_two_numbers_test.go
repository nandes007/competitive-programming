package tests

import (
	"testing"

	"github.com/nandes007/leetcode/golang/exercise"
	"github.com/nandes007/leetcode/golang/model"
	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	// Helper function to convert an array to a ListNode
	arrayToList := func(arr []int) *model.ListNode {
		dummy := &model.ListNode{}
		current := dummy
		for _, val := range arr {
			current.Next = &model.ListNode{Val: val}
			current = current.Next
		}
		return dummy.Next
	}

	tests := []struct {
		l1     *model.ListNode
		l2     *model.ListNode
		result *model.ListNode
	}{
		{
			arrayToList([]int{2, 4, 3}),
			arrayToList([]int{5, 6, 4}),
			arrayToList([]int{7, 0, 8}),
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := exercise.AddTwoNumbers(test.l1, test.l2)
		assert.Equal(t, test.result, result, "For l1=%v, l2=%v", test.l1, test.l2)
	}
}

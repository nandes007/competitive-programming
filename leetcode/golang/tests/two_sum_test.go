package tests

import (
	"testing"

	"github.com/nandes007/leetcode/golang/exercise"
	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := exercise.TwoSum(test.nums, test.target)
		assert.Equal(t, test.result, result, "For nums=%v, target=%d", test.nums, test.target)
	}
}

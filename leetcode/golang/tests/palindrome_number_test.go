package tests

import (
	"testing"

	"github.com/nandes007/leetcode/golang/exercise"
	"github.com/stretchr/testify/assert"
)

func TestPlindromeNumber(t *testing.T) {
	assert.False(t, exercise.IsPalindrome(10))
	assert.True(t, exercise.IsPalindrome(121))
	assert.True(t, exercise.IsPalindrome(111))

	assert.False(t, exercise.IsPalindromeV2(10))
	assert.True(t, exercise.IsPalindromeV2(121))
	assert.True(t, exercise.IsPalindromeV2(111))
}

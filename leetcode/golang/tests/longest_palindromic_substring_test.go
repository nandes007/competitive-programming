package tests

import (
	"testing"

	"github.com/nandes007/leetcode/golang/exercise"
	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	s1 := "babad"
	s2 := "cbbd"
	assert.Equal(t, "bab", exercise.LongestPalindrome(s1))
	assert.Equal(t, "bb", exercise.LongestPalindrome(s2))
}

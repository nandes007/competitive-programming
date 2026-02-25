package tests

import (
	"testing"

	"github.com/nandes007/leetcode/golang/exercise"
	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s1 := "abcabcbb"
	s2 := "bbbbb"
	s3 := "pwwkew"

	assert.Equal(t, 3, exercise.LengthOfLongestSubstring(s1))
	assert.Equal(t, 1, exercise.LengthOfLongestSubstring(s2))
	assert.Equal(t, 3, exercise.LengthOfLongestSubstring(s3))
}

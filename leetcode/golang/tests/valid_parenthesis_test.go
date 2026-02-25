package tests

import (
	"testing"

	"github.com/nandes007/leetcode/golang/exercise"
	"github.com/stretchr/testify/assert"
)

func TestIsValidParenthesis(t *testing.T) {
	s1 := "()[]{}"
	s2 := "(]"
	s3 := "()"
	assert.True(t, exercise.IsValid(s1))
	assert.False(t, exercise.IsValid(s2))
	assert.True(t, exercise.IsValid(s3))
}

func TestIsValidParenthesisV2(t *testing.T) {
	s1 := "()[]{}"
	s2 := "(]"
	s3 := "()"
	assert.True(t, exercise.IsValid(s1))
	assert.False(t, exercise.IsValid(s2))
	assert.True(t, exercise.IsValid(s3))
}

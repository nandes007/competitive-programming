package tests

import (
	"testing"

	"github.com/nandes007/leetcode/golang/exercise"
	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	strsOne := []string{"flower", "flow", "flight"}
	strsTwo := []string{"dog", "racecar", "car"}

	assert.Equal(t, "fl", exercise.LongestCommonPrefix(strsOne))
	assert.Equal(t, "", exercise.LongestCommonPrefix(strsTwo))
}

package tests

import (
	"testing"

	"github.com/nandes007/leetcode/golang/exercise"
	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	assert.Equal(t, 1994, exercise.RomanToInt("MCMXCIV"))
	assert.Equal(t, 58, exercise.RomanToInt("LVIII"))
	assert.Equal(t, 3, exercise.RomanToInt("III"))
}

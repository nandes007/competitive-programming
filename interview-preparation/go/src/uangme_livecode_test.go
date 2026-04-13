package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRaceCondition(t *testing.T) {
	RaceCondition()
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, 20, multiply(10))
}

func TestLargestNumber(t *testing.T) {
	assert.Equal(t, 6, largestNumber([]int{1, 2, 3}, []int{4, 5, 6}))
}

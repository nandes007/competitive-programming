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

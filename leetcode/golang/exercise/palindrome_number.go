package exercise

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	xString := strconv.Itoa(x)
	var temp string

	for i := len(xString) - 1; i >= 0; i-- {
		temp = temp + string([]rune(xString)[i])
	}

	return xString == temp
}

func IsPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x = x / 10
	}

	return original == reversed
}

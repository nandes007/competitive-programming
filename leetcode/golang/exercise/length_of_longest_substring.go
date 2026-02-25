package exercise

func LengthOfLongestSubstring(s string) int {
	n := len(s)
	res := 0
	left, right := 0, 0
	seen := make(map[rune]int)

	for right < n {
		currentChar := rune(s[right])
		if val, ok := seen[currentChar]; ok {
			left = max(left, val+1)
		}
		seen[currentChar] = right
		res = max(res, right-left+1)
		right++
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

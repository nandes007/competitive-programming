package exercise

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		current := strs[i]
		j := 0

		for j < len(prefix) && j < len(current) && prefix[j] == current[j] {
			j++
		}

		prefix = prefix[:j]

		if prefix == "" {
			break
		}
	}

	return prefix
}

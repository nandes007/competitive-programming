package exercise

func IsValid(s string) bool {
	parenthesis := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	var openParenthesis string
	var closeParenthesis string
	var isValid bool
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			openParenthesis = string([]rune(s)[i])
			_, exists := parenthesis[openParenthesis]
			if !exists {
				isValid = exists
				break
			} else {
				closeParenthesis = parenthesis[openParenthesis]
			}
		} else {
			_, exists := parenthesis[openParenthesis]
			if exists {
				if string([]rune(s)[i]) == closeParenthesis {
					isValid = exists
				} else {
					isValid = false
					break
				}
			} else {
				isValid = false
				break
			}
		}
	}

	return isValid
}

func IsValidV2(s string) bool {
	parenthesis := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := make([]rune, 0)

	for _, char := range s {
		if closing, exists := parenthesis[char]; exists {
			stack = append(stack, closing)
		} else {
			if len(stack) == 0 || char != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

package main

// 20题： https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	tmpCache := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		tmp := s[i]

		v, ok := tmpCache[tmp]
		if !ok {
			stack = append(stack, tmp)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func isValid(s string) bool {
	matchingBrackets := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
		
			if len(stack) == 0 || stack[len(stack)-1] != matchingBrackets[char] {
				return false
			}
			
			stack = stack[:len(stack)-1]
		}
	}

	
	return len(stack) == 0
}
package soal

func IsValidBrackets(s string) bool {
	stack := []rune{}
	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		// opening bracket dapat selalu dilakukan secara unlimited untuk jenis bracket
		case '(', '{', '[':
			stack = append(stack, char)
			// fmt.Println(string(stack))
		case ')', '}', ']':
			// closing bracket dapat dilakukan dengan kondisi
			// sudah ada minimal 1 opening bracket sebelumnya
			// bracket yang sedang berjalan harus match dengan bracket opening sebelumnya
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return false
			}
			// jika bracket match, maka hapus bracket yang sudah closing
			stack = stack[:len(stack)-1]
			// fmt.Println(string(stack))
		}
	}

	return len(stack) == 0
}

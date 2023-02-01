package exemplo4

func palindromeFromEnd(str string) bool {
	runes := []rune(str)
	length := len(runes)
	for i := 0; i < length; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

func palindromeReverse(str string) bool {
	strReversed := reverse(str)
	return strReversed == str
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

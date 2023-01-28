package exemplo4

import "strings"

func findIndexByEduardo(str string, term string) int {
	strRunes := []rune(str)
	termRune := []rune(term)
	for i := 0; i < len(strRunes); i++ {
		if strRunes[i] == termRune[0] {
			return i
		}
	}
	return -1
}

func findIndexByGolang(str string, term string) int {
	return strings.Index(str, term)
}

package exemplo4

import "testing"

func BenchmarkPalindromeReverse_Hannah(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeReverse("hannah")
	}
}

func BenchmarkPalindromeFromEnd_Hannah(b *testing.B) {
	for n := 0; n < b.N; n++ {
		palindromeFromEnd("hannah")
	}

}

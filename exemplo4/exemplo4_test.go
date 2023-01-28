package exemplo4

import "testing"

func BenchmarkFindIndexByEduardo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findIndexByEduardo("Golang", "ang")
	}
	b.ReportAllocs()
}

func BenchmarkFindIndexByGolang(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findIndexByGolang("Golang", "ang")
	}
	b.ReportAllocs()

}

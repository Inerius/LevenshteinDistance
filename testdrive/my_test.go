package testdrive

import (
	"testing"

	"../levenshteindistance"
)

// BenchmarkGetStrings - benchmark for GetStrings function
func BenchmarkGetStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		levenshteindistance.GetStrings("../main/input.txt")
	}
}

// BenchmarkMakePairs - benchmark for MakePairs function
func BenchmarkMakePairs(b *testing.B) {
	keyword, strlst := levenshteindistance.GetStrings("../main/input.txt")
	for i := 0; i < b.N; i++ {
		levenshteindistance.MakePairs(keyword, strlst)
	}
}

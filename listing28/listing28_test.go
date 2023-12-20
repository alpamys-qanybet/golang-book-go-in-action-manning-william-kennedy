package listing28_test

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	n := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", n)
	}
}

// go test -v -run="none" -bench="BenchmarkSprintf"
// go test -v -run="none" -bench="BenchmarkSprintf" -benchtime="3s"

func BenchmarkFormat(b *testing.B) {
	n := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(n, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	n := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(n)
	}
}

// go test -v -run="none" -bench=.
// go test -v -run="none" -bench=. -benchtime="3s"
// go test -v -run="none" -bench=. -benchmem

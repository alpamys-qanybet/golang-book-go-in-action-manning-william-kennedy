package listing28_test

import (
  "fmt"
  // "strconv"
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

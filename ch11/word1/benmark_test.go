package word

import "testing"

// Benchmark
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		IsPalindrome("amaama")
	}
}

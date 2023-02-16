package main

import "testing"

func TestPlaceholder(t *testing.T) {
	x := 10
	got := placeholder(x)
	if got != 10 {
		t.Errorf("Expected: %v\t Got: %v", 10, got)
	}
}

func BenchmarkPlaceholder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		placeholder(10)
	}
}

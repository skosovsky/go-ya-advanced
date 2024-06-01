package main

import "testing"

func BenchmarkConcatPlus(b *testing.B) {
	values := []string{"test1", "test2"}

	for i := 0; i < b.N; i++ {
		ConcatPlus(values)
	}
}

func BenchmarkConcatSprintf(b *testing.B) {
	values := []string{"test1", "test2"}

	for i := 0; i < b.N; i++ {
		ConcatSprintf(values)
	}
}

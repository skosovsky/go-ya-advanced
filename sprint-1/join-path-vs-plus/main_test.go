package main

import "testing"

func BenchmarkPathJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlJoinPath()
	}
}

func BenchmarkPathPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlJoinPlus()
	}
}

func BenchmarkPathSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlJoinSprintf()
	}
}

package main

import "testing"

func BenchmarkInsert(b *testing.B) {
	zero := New(20)

	for i := 0; i < b.N; i++ {
		zero.Insert()
	}
}

func BenchmarkInsertIf(b *testing.B) {
	zero := New(1)

	for i := 0; i < b.N; i++ {
		zero.InsertIf()
	}
}

// Benchmark passing a [10000]int struct by Value vs Pointer. Measure the speed difference.

package main

import "testing"

type Heavy struct {
	Data [10000]int
}

var result int

func byValue(h Heavy) {
	result = h.Data[9999]
}

func byPointer(h *Heavy) {
	result = h.Data[9999]
}

func BenchmarkByValue(b *testing.B) {
	h := Heavy{}

	for i := 0; i < b.N; i++ {
		byValue(h)
	}
}

func BenchmarkByPointer(b *testing.B) {
	h := Heavy{}

	for i := 0; i < b.N; i++ {
		byPointer(&h)
	}
}

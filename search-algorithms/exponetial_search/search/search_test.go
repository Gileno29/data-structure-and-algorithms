package search

import (
	"testing"
)

// Cria um array ordenado para os testes
func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i * 2
	}
	return arr
}

// Benchmark do Binary Search (Elemento no Início)
func BenchmarkBinarySearchEarly(b *testing.B) {
	size := 1000000
	arr := generateSortedArray(size)
	target := 10 // Perto do início

	for i := 0; i < b.N; i++ {
		BinarySearch(arr, target, 0, size-1)
	}
}

// Benchmark do Exponential Search (Elemento no Início)
func BenchmarkExponentialSearchEarly(b *testing.B) {
	size := 1000000
	arr := generateSortedArray(size)
	target := 10 // Perto do início

	for i := 0; i < b.N; i++ {
		ExponentialSearch(arr, target)
	}
}

// Benchmark do Binary Search (Elemento no Fim)
func BenchmarkBinarySearchLate(b *testing.B) {
	size := 1000000
	arr := generateSortedArray(size)
	target := 1999990 // Perto do fim

	for i := 0; i < b.N; i++ {
		BinarySearch(arr, target, 0, size-1)
	}
}

// Benchmark do Exponential Search (Elemento no Fim)
func BenchmarkExponentialSearchLate(b *testing.B) {
	size := 1000000
	arr := generateSortedArray(size)
	target := 1999990 // Perto do fim

	for i := 0; i < b.N; i++ {
		ExponentialSearch(arr, target)
	}
}

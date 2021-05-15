package sort

import (
	"testing"
)

func TestBubbleSortOrderDESC(t *testing.T) {
	// Arrange
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Act
	BubbleSort(elements)

	// Assert
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}
	if elements[len(elements)-1] != 0 {
		t.Error("last elements should be 0")
	}
}

func TestSortIncreasingOrder(t *testing.T) {
	// Arrange
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Act
	Sort(elements)

	// Assert
	if elements[0] != 0 {
		t.Error("first element should be 9")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last elements should be 0")
	}
}

// test helper
func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n; i > 0; i-- {
		result[j] = i
		j++
	}

	return result
}

func BenchmarkBubbleSort(b *testing.B) {
	// Arrange
	elements := getElements(10000)

	// Act
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	// Arrange
	elements := getElements(10000)

	// Act
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

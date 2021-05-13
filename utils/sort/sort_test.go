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

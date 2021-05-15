package services

import (
	"testing"
)

func TestConstants(t *testing.T) {
	if privateConst != "private" {
		t.Error("privateConst should be 'private'")
	}
}

func TestSort(t *testing.T) {
	// Arrange
	elements := getElements(10)

	// Act
	Sort(elements)

	// Assert
	if elements[0] != 10 {
		t.Error("first element should be 10")
	}
	if elements[len(elements)-1] != 1 {
		t.Error("last elements should be 1")
	}
}

func TestBubbleSort(t *testing.T) {
	// Arrange
	elements := getElements(10001)

	// Act
	Sort(elements)

	// Assert
	if elements[0] != 1 {
		t.Error("first element should be 1")
	}
	if elements[len(elements)-1] != 10001 {
		t.Error("last elements should be 10001")
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

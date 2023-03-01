package tests

import (
	"start/algo"
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSorting(t *testing.T) {
	arr := []int{5, 4, 1, 3, 2}
	sortedArray := algo.InsertionSort(arr)
	expectedArray := []int{1, 2, 3, 4, 5}

	if !equal(sortedArray, expectedArray) {
		t.Errorf("Expected array %v, got %v", expectedArray, sortedArray)
	}
}

func TestSortingReverseOrder(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	sortedArray := algo.InsertionSort(arr)
	expectedArray := []int{1, 2, 3, 4, 5}

	if !equal(sortedArray, expectedArray) {
		t.Errorf("Expected array %v, got %v", expectedArray, sortedArray)
	}
}

func TestSortingRandomOrder(t *testing.T) {
	arr := []int{3, 5, 2, 1, 4}
	sortedArray := algo.InsertionSort(arr)
	expectedArray := []int{1, 2, 3, 4, 5}

	if !equal(sortedArray, expectedArray) {
		t.Errorf("Expected array %v, got %v", expectedArray, sortedArray)
	}
}

func TestSortingNegativeValues(t *testing.T) {
	arr := []int{-1, -5, 2, -3, 4}
	sortedArray := algo.InsertionSort(arr)
	expectedArray := []int{-5, -3, -1, 2, 4}

	if !equal(sortedArray, expectedArray) {
		t.Errorf("Expected array %v, got %v", expectedArray, sortedArray)
	}
}

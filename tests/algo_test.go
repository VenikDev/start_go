package tests

import (
	"github.com/stretchr/testify/assert"
	"start/algo"
	"testing"
)

func TestAdd(t *testing.T) {
	// Test 1
	result := algo.Add(2, 3)
	if result != 5 {
		t.Errorf("Expected result to be 5, got %d", result)
	}

	// Test 2
	result = algo.Add(-2, 4)
	if result != 2 {
		t.Errorf("Expected result to be 2, got %d", result)
	}
}

func TestFindMaximum(t *testing.T) {
	// test case 1
	numbers := []int{1, 2, 3}
	expectedMax := 3
	actualMax := algo.FindMaximum(numbers)

	if actualMax != expectedMax {
		t.Errorf("Expected %d, but got %d", expectedMax, actualMax)
	}

	// test case 2
	numbers = []int{-1, 0, 10}
	expectedMax = 10
	actualMax = algo.FindMaximum(numbers)

	if actualMax != expectedMax {
		t.Errorf("Expected %d, but got %d", expectedMax, actualMax)
	}
}

func TestTwoLargest(t *testing.T) {
	// Test empty array
	inputArray := []int{}
	_, _ = algo.TwoLargest(inputArray)
	assert.Equal(t, 0, 0, "No elements should return 0")

	// Test single element array
	inputArray = []int{1}
	largest, secondLargest := algo.TwoLargest(inputArray)
	assert.Equal(t, largest, secondLargest, "Single element array should return same element for both")

	// Test multiple elements array
	inputArray = []int{1, 2, 3, 4, 5}
	largest, secondLargest = algo.TwoLargest(inputArray)
	assert.Equal(t, largest, 5, "Largest element should be 5")
	assert.Equal(t, secondLargest, 4, "Second largest element should be 4")
}

func TestMinMaxNegatives(t *testing.T) {
	// Тестируем функцию MinMax с аргументами:
	// []int{-2, -5, -1, -4, -3}
	min, max := algo.MinMax([]int{-2, -5, -1, -4, -3})

	// Проверяем, что min и max равны -5 и -1 соответственно
	if min != -5 && max != -1 {
		t.Errorf("Incorrect output of MinMax(), got %d and %d, want -5 and -1", min, max)
	}
}

func TestMinMaxSingleNumber(t *testing.T) {
	// Тестируем функцию MinMax с аргументами:
	// []int{7}
	min, max := algo.MinMax([]int{7})

	// Проверяем, что min и max равны 7
	if min != 7 && max != 7 {
		t.Errorf("Incorrect output of MinMax(), got %d and %d, want 7 and 7", min, max)
	}
}

func TestMinMaxEmptySlice(t *testing.T) {
	// Тестируем функцию MinMax с аргументами:
	// []int{}
	min, max := algo.MinMax([]int{})

	// Проверяем, что min и max равны -1
	if min != -1 && max != -1 {
		t.Errorf("Incorrect output of MinMax(), got %d and %d, want -1 and -1", min, max)
	}
}

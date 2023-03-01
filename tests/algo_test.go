package tests

import (
	"github.com/stretchr/testify/assert"
	"reflect"
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

func TestSum(t *testing.T) {
	// Test cases for Sum function
	testCases := []struct {
		name string
		arr  []int
		sum  int
	}{
		{
			name: "Simple Addition",
			arr:  []int{1, 2},
			sum:  3,
		},
		{
			name: "Adding Empty List",
			arr:  []int{},
			sum:  0,
		},
		{
			name: "Adding Negatives",
			arr:  []int{-1, -2, -3},
			sum:  -6,
		},
		{
			name: "Large List",
			arr:  []int{10, 20, 30, 40, 50},
			sum:  150,
		},
		{
			name: "Empty arr",
			arr:  []int{},
			sum:  0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sum := algo.Sum(tc.arr)
			if sum != tc.sum {
				t.Errorf("expected sum of array %v to be %d but got %d", tc.arr, tc.sum, sum)
			}
		})
	}
}

// testing reverse with different inputs
func TestReverse(t *testing.T) {
	testCases := []struct {
		in  []int
		out []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{7, 6, 5, 4, 3, 2, 1}},
		{[]int{2, 4, 1}, []int{1, 4, 2}},
		{[]int{-2, -4, -1}, []int{-1, -4, -2}},
		{[]int{}, []int{}},
	}
	//run each test case
	for _, c := range testCases {
		algo.Reverse(c.in)
		//compare result to expected output
		if !reflect.DeepEqual(c.in, c.out) {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, c.in, c.out)
		}
	}
}

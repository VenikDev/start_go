package tests

import (
	"reflect"
	"start/bytes"
	"testing"
)

// TestShiftLeft tests that ShiftLeft correctly shifts a byte slice to the left
func TestShiftLeft(t *testing.T) {
	original := []byte{1, 2, 3, 4}
	expected := []byte{2, 3, 4, 0}

	bytes.ShiftLeft(original)

	for i := range original {
		if original[i] != expected[i] {
			t.Errorf("index %d: expected %d, got %d", i, expected[i], original[i])
		}
	}
}

// TestShiftLeftEmpty tests that ShiftLeft correctly handles an empty byte slice
func TestShiftLeftEmpty(t *testing.T) {
	original := []byte{}

	bytes.ShiftLeft(original)

	if len(original) != 0 {
		t.Error("expected an empty slice, got a non-empty slice")
	}
}

// TestShiftLeftRemaining tests that ShiftLeft correctly sets the remaining byte to 0
func TestShiftLeftRemaining(t *testing.T) {
	original := []byte{1, 2}
	expected := []byte{2, 0}

	bytes.ShiftLeft(original)

	if original[1] != expected[1] {
		t.Errorf("index 1: expected %d, got %d", expected[1], original[1])
	}
}

func TestShiftRight1(t *testing.T) {
	bs := []byte{1, 2, 3, 4}
	bytes.ShiftRight(bs)
	expected := []byte{0, 1, 2, 3}
	if !reflect.DeepEqual(bs, expected) {
		t.Errorf("Expected %v, got %v", expected, bs)
	}
}

func TestShiftRight2(t *testing.T) {
	bs := []byte{1, 2, 3}
	bytes.ShiftRight(bs)
	expected := []byte{0, 1, 2}
	if !reflect.DeepEqual(bs, expected) {
		t.Errorf("Expected %v, got %v", expected, bs)
	}
}

func TestShiftRight3(t *testing.T) {
	bs := []byte{1, 2}
	bytes.ShiftRight(bs)
	expected := []byte{0, 1}
	if !reflect.DeepEqual(bs, expected) {
		t.Errorf("Expected %v, got %v", expected, bs)
	}
}

func TestShiftRight4(t *testing.T) {
	bs := []byte{1}
	bytes.ShiftRight(bs)
	expected := []byte{0}
	if !reflect.DeepEqual(bs, expected) {
		t.Errorf("Expected %v, got %v", expected, bs)
	}
}

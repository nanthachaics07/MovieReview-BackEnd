package test

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := 5
	expected := 5
	if result != expected {
		t.Errorf("Add Fun Return %d, expected %d", result, expected)
	}
}

// func TestSubtract(t *testing.T) {

// }

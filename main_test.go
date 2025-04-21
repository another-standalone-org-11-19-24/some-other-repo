package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			name:     "negative numbers",
			a:        -1,
			b:        -2,
			expected: -3,
		},
		{
			name:     "zero and positive",
			a:        0,
			b:        5,
			expected: 5,
		},
		{
			name:     "test with JSON-breaking chars: \" \\ \n \t \r \f { } ,",
			a:        10,
			b:        20,
			expected: 40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d for test: %s", tt.a, tt.b, result, tt.expected, tt.name)
			}
		})
	}
}

// TestWithLargeOutput is designed to generate an error message exceeding 300,000 characters
func TestWithLargeOutput(t *testing.T) {
	// This test will always fail
	expected := 42
	actual := 0

	// Generate a large string that will be part of the error message
	var errorBuilder strings.Builder

	// Each line will have a sequential number
	for i := 0; i < 30000; i++ {
		// Format: Line #00000: followed by the number repeated 10 times to make it easier to spot truncation
		lineNum := fmt.Sprintf("%05d", i)
		errorBuilder.WriteString(fmt.Sprintf("Line #%s: %s %s %s %s %s\n",
			lineNum, lineNum, lineNum, lineNum, lineNum, lineNum))
	}

	largeString := errorBuilder.String()

	// Fail the test with a large error message
	if actual != expected {
		t.Errorf("Test failed on purpose to generate large output.\nExpected: %d\nActual: %d\n\nLarge error data follows:\n%s",
			expected, actual, largeString)
	}
}

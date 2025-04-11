package main

import "testing"

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

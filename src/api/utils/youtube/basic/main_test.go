package main

import "testing"

func TestAddNumber(t *testing.T) {
	if addNumber(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableAddNumber(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{9999, 10001},
	}

	for _, test := range tests {
		if output := addNumber(test.input); output != test.expected {
			t.Error("Test Failes: {} inputted, {} expected", test.input, test.expected, output)
		}
	}
}

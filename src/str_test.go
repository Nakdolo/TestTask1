package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HIIII", "IIIIH"},
		{"Golang", "gnaloG"},
		{"", ""},
		{"12345", "54321"},
		{"ss", "ss"},
		{"Intern", "nretnI"},
		{"", ""},
		{"#!3421431_3", "3_1341243!#"},
	}

	for _, test := range tests {
		result := ReverseString(&test.input)

		if result != test.expected {
			t.Errorf("%v != %v, expected %v", test.input, result, test.expected)
		}
	}
}



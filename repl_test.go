package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	// table-driven test
	cases := []struct {
	input    string
	expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "How    you doing?",
			expected: []string{"how", "you", "doing?"},
		},
		{
			input:    "  leading and trailing spaces  ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("cleanInput(%q) = %v; expected %v", c.input, actual, c.expected)
		}
	}

}
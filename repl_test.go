package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	}{
			{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input: "Charmander Bulbasaur PIKACHU",
		expected: []string{"charmander", "bulbasaur", "pikachu"},
	},
	}

	// run the cases
	for _, c := range cases {
		actual := cleanInput(c.input)
		// fmt.Printf("Actual is: %s", actual)
		// fmt.Printf("expected is: %s", c.expected)
		if len(actual) != len(c.expected) {
			t.Errorf("FAILED: Expected %v, Got %v", len(c.expected), len(actual))
			t.Fail()
		}
			// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]

		if word != expectedWord {
			t.Errorf("FAILED: Expected %s: Got %s", expectedWord, word)
			t.Fail()
		}
	}
	}
}

// package main

// import (
// 	"testing"
// )

// func TestCleanInput(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		expected []string
// 	}{
// 		{
// 			input:    "  hello  world  ",
// 			expected: []string{"hello", "world"},
// 		},
// 		{
// 			input:    "Charmander Bulbasaur PIKACHU",
// 			expected: []string{"charmander", "bulbasaur", "pikachu"},
// 		},
// 		{
// 			input:    "   A   B  C ",
// 			expected: []string{"a", "b", "c"},
// 		},
// 		{
// 			input:    "",
// 			expected: []string{},
// 		},
// 		{
// 			input:    "   ",
// 			expected: []string{},
// 		},
// 	}

// 	for _, c := range cases {
// 		actual := cleanInput(c.input)

// 		if len(actual) != len(c.expected) {
// 			t.Errorf("For input %q, expected length %d, got %d. Full output: %v",
// 				c.input, len(c.expected), len(actual), actual)
// 			continue
// 		}

// 		for i := range actual {
// 			if actual[i] != c.expected[i] {
// 				t.Errorf("For input %q, expected word %q at index %d, got %q",
// 					c.input, c.expected[i], i, actual[i])
// 			}
// 		}
// 	}
// }

package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}
	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("the lengths are not equal: %v vs %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectWord := cs.expected[i]
			if actualWord != expectWord {
				t.Errorf("%v does not equal %v", actualWord, expectWord)
			}
		}
	}
}
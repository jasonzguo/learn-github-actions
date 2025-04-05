package utils

import (
	"testing"
)

type TestCase struct {
	name   string
	inputs []string
	output string
}

func TestAddStrings(t *testing.T) {
	testCases := []TestCase{
		{
			name:   "no carry",
			inputs: []string{"1", "2"},
			output: "3",
		},
		{
			name:   "no carry - unequal length - len(a) > len(b)",
			inputs: []string{"10", "9"},
			output: "19",
		},
		{
			name:   "no carry - unequal length - len(b) > len(a)",
			inputs: []string{"9", "10"},
			output: "19",
		},
		{
			name:   "carry",
			inputs: []string{"1", "9"},
			output: "10",
		},
		{
			name:   "carry - with more than 1 digits",
			inputs: []string{"25", "25"},
			output: "50",
		},
		{
			name:   "carry - twice",
			inputs: []string{"257", "253"},
			output: "510",
		},
		{
			name:   "invalid inputs",
			inputs: []string{"abc", "253"},
			output: "0",
		},
	}

	for _, testCase := range testCases {
		actualOutput, _ := AddStrings(testCase.inputs[0], testCase.inputs[1])
		if testCase.output != actualOutput {
			t.Errorf("got %s, want %s", actualOutput, testCase.output)
		}
	}
}

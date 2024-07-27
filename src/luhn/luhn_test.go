package luhn

import "testing"

func TestValid(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Valid(tc.input); actual != tc.expected {
				t.Fatalf("Valid(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}
func BenchmarkValid(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Valid(tc.input)
		}
	}
}

var testCases = []struct {
	description string
	input       string
	expected    bool
}{
	{
		description: "single digit strings can not be valid",
		input:       "1",
		expected:    false,
	},
	{
		description: "a single zero is invalid",
		input:       "0",
		expected:    false,
	},
	{
		description: "a simple valid SIN that remains valid if reversed",
		input:       "059",
		expected:    true,
	},
	{
		description: "a simple valid SIN that becomes invalid if reversed",
		input:       "59",
		expected:    true,
	},
	{
		description: "a valid Canadian SIN",
		input:       "055 444 285",
		expected:    true,
	},
	{
		description: "invalid Canadian SIN",
		input:       "055 444 286",
		expected:    false,
	},
	{
		description: "invalid credit card",
		input:       "8273 1232 7352 0569",
		expected:    false,
	},
	{
		description: "invalid long number with an even remainder",
		input:       "1 2345 6789 1234 5678 9012",
		expected:    false,
	},
	{
		description: "invalid long number with a remainder divisible by 5",
		input:       "1 2345 6789 1234 5678 9013",
		expected:    false,
	},
	{
		description: "valid number with an even number of digits",
		input:       "095 245 88",
		expected:    true,
	},
	{
		description: "valid number with an odd number of spaces",
		input:       "234 567 891 234",
		expected:    true,
	},
	{
		description: "valid strings with a non-digit added at the end become invalid",
		input:       "059a",
		expected:    false,
	},
	{
		description: "valid strings with punctuation included become invalid",
		input:       "055-444-285",
		expected:    false,
	},
	{
		description: "valid strings with symbols included become invalid",
		input:       "055# 444$ 285",
		expected:    false,
	},
	{
		description: "single zero with space is invalid",
		input:       " 0",
		expected:    false,
	},
	{
		description: "more than a single zero is valid",
		input:       "0000 0",
		expected:    true,
	},
	{
		description: "input digit 9 is correctly converted to output digit 9",
		input:       "091",
		expected:    true,
	},
	{
		description: "very long input is valid",
		input:       "9999999999 9999999999 9999999999 9999999999",
		expected:    true,
	},
	{
		description: "valid luhn with an odd number of digits and non zero first digit",
		input:       "109",
		expected:    true,
	},
	{
		description: "using ascii value for non-doubled non-digit isn't allowed",
		input:       "055b 444 285",
		expected:    false,
	},
	{
		description: "using ascii value for doubled non-digit isn't allowed",
		input:       ":9",
		expected:    false,
	},
	{
		description: "non-numeric, non-space char in the middle with a sum that's divisible by 10 isn't allowed",
		input:       "59%59",
		expected:    false,
	},
}

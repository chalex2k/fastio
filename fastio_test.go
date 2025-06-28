package main

import (
	"bytes"
	"testing"
)

func TestPrintSingleArgument(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected string
	}{
		{"Positive number", 42, "42\n"},
		{"Zero", 0, "0\n"},
		{"Negative number", -3, "-3\n"},
		{"String", "line", "line\n"},
		{"Empty line", "", "\n"},
		{"Char", 'A', "A\n"},
		{"True", true, "true\n"},
		{"False", false, "false\n"},
		{"Uint64", uint64(30_000_000_000), "30000000000\n"},
		{"Float number", 1.333555, "1.333555\n"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			out := new(bytes.Buffer)
			outputStream = out

			print(testCase.input)

			if out.String() != testCase.expected {
				t.Errorf("print(%d) = %q, want %q", testCase.input, out, testCase.expected)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	testCases := []struct {
		name     string
		input    []any
		expected string
	}{
		{"Multiple ints", []any{1, -4, 0, 155}, "1 -4 0 155\n"},
		{"Multiple different", []any{1, "str", 'x', '0', 1.5}, "1 str x 0 1.5\n"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			out := new(bytes.Buffer)
			outputStream = out

			print(testCase.input...)

			if out.String() != testCase.expected {
				t.Errorf("print(%d) = %q, want %q", testCase.input, out, testCase.expected)
			}
		})
	}
}

package main

import (
	"bufio"
	"bytes"
	"strings"
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

func TestPrintMultipleArguments(t *testing.T) {
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

func TestPrintSlice(t *testing.T) {
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

			print(testCase.input)

			if out.String() != testCase.expected {
				t.Errorf("print(%d) = %q, want %q", testCase.input, out, testCase.expected)
			}
		})
	}
}

func TestPrintSliceWithSep(t *testing.T) {
	testCases := []struct {
		name     string
		input    []any
		sep      string
		expected string
	}{
		{"Multiple ints", []any{1, -4, 0, 155}, " ", "1 -4 0 155\n"},
		{"Multiple ints", []any{1, -4, 0, 155}, "-", "1--4-0-155\n"},
		{"Multiple different", []any{1, "str", 'x', '0', 1.5}, "", "1strx01.5\n"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			out := new(bytes.Buffer)
			outputStream = out

			print(testCase.input, testCase.sep)

			if out.String() != testCase.expected {
				t.Errorf("print(%d) = %q, want %q", testCase.input, out, testCase.expected)
			}
		})
	}
}

func TestInputInt(t *testing.T) {
	testCases := []struct {
		name     string
		stdin    string
		expected int
	}{
		{"Positive", "5", 5},
		{"Positive with new line", "5\n", 5},
		{"Positive with new line and spaces", " 5 \n", 5},
		{"Negative", "-999", -999},
		{"Zero", "0", 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			in := strings.NewReader(testCase.stdin)
			inputStream = in
			scanner = bufio.NewScanner(inputStream)

			out := ii()

			if out != testCase.expected {
				t.Errorf("%v: expected %v, found %v", testCase.stdin, testCase.expected, out)
			}
		})
	}
}

func TestListInputInt(t *testing.T) {
	testCases := []struct {
		name     string
		stdin    string
		expected []int
	}{
		{"Usual list", "5 0 -4 1111111", []int{5, 0, -4, 1111111}},
		{"List with new line", "5 0 -4 1111111\n", []int{5, 0, -4, 1111111}},
		{"Positive with new line and spaces", " 5  0 -4 1111111  \n", []int{5, 0, -4, 1111111}},
		{"One item in list", "-999", []int{-999}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			in := strings.NewReader(testCase.stdin)
			inputStream = in
			scanner = bufio.NewScanner(inputStream)

			out := lii()
			if len(out) != len(testCase.expected) {
				t.Errorf("%v: expected %v, got %v", testCase.stdin, testCase.expected, out)
			}

			for i := range out {
				if out[i] != testCase.expected[i] {
					t.Errorf("%v: expected %v, found %v", testCase.stdin, testCase.expected, out)
				}
			}
		})
	}
}

func TestInput2Ints(t *testing.T) {
	testCases := []struct {
		name      string
		stdin     string
		expected1 int
		expected2 int
	}{
		{"Both positive", "3 6", 3, 6},
		{"Both negative", "-3000 -1", -3000, -1},
		{"Both zeros", "0 0", 0, 0},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			in := strings.NewReader(testCase.stdin)
			inputStream = in
			scanner = bufio.NewScanner(inputStream)

			out1, out2 := i2()

			if out1 != testCase.expected1 {
				t.Errorf("first number error. %v: expected %v, got %v",
					testCase.stdin, testCase.expected1, out1)
			}

			if out2 != testCase.expected2 {
				t.Errorf("second number error. %v: expected %v, got %v",
					testCase.stdin, testCase.expected2, out2)
			}
		})
	}
}

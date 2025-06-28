package main

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
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

// Генерируем тестовые данные один раз перед всеми тестами
var (
	stringArrays [500][]any
	numberArrays [500][]any
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// Генерация 500 массивов строк
	for i := 0; i < 500; i++ {
		arr := make([]any, 1000)
		for j := 0; j < 1000; j++ {
			arr[j] = randomString(10)
		}
		stringArrays[i] = arr
	}

	// Генерация 500 массивов чисел
	for i := 0; i < 500; i++ {
		arr := make([]any, 1000)
		for j := 0; j < 1000; j++ {
			arr[j] = rand.Intn(10000)
		}
		numberArrays[i] = arr
	}
}

// Вспомогательная функция для генерации случайных строк
func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Бенчмарк для строковых массивов
func BenchmarkPrintStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Берем массив по модулю, чтобы не выйти за границы
		arr := stringArrays[i%500]
		print(arr...)
	}
}

// Бенчмарк для числовых массивов
func BenchmarkPrintNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := numberArrays[i%500]
		print(arr...)
	}
}

// Комбинированный бенчмарк
func BenchmarkPrintMixed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Чередуем строковые и числовые массивы
		if i%2 == 0 {
			print(stringArrays[i%500]...)
		} else {
			print(numberArrays[i%500]...)
		}
	}
}

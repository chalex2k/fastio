package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	x := ii()
	print(x)

	arr := lii()
	print(arr)
}

//nolint:gochecknoglobals
var (
	outputStream io.Writer
	inputStream  io.Reader
	scanner      *bufio.Scanner
)

func initStdStreams() {
	outputStream = os.Stdout
	inputStream = os.Stdin
	scanner = bufio.NewScanner(inputStream)
}

func init() {
	initStdStreams()
}

const sepCaseArgsLen = 2

//nolint:revive,predeclared
func print(params ...any) {
	var res []string

	sep := " "

	if ok, collection, sep2 := isSepCase(params...); ok {
		res = processArgs(collection)
		sep = sep2
	} else {
		for _, v := range params {
			res = append(res, processArgs(v)...)
		}
	}

	_, err := fmt.Fprintln(outputStream, strings.Join(res, sep))
	if err != nil {
		panic(err)
	}
}

func isSepCase(args ...any) (bool, any, string) {
	if len(args) != sepCaseArgsLen {
		return false, nil, ""
	}

	first := reflect.ValueOf(args[0])
	second := reflect.ValueOf(args[1])

	if (first.Kind() == reflect.Slice || first.Kind() == reflect.Array) &&
		second.Kind() == reflect.String {
		sep, ok := args[1].(string)
		if !ok {
			panic(fmt.Sprintf("%v not a string", sep))
		}

		return true, args[0], sep
	}

	return false, nil, ""
}

func processArgs(args any) []string {
	var res []string

	val := reflect.ValueOf(args)

	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		for i := range val.Len() {
			res = append(res, processArgs(val.Index(i).Interface())...)
		}
	} else {
		switch val := args.(type) {
		case rune:
			res = append(res, fmt.Sprintf("%c", val))
		default:
			res = append(res, fmt.Sprintf("%v", val))
		}
	}

	return res
}

func input() string {
	if scanner.Scan() {
		return scanner.Text()
	} else if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Fatal("not scan and not err")

	return ""
}

func ii() int {
	s := input()
	s = strings.TrimSpace(s)
	num, _ := strconv.Atoi(s)

	return num
}

func lii() []int {
	parts := strings.Fields(input())
	numbers := make([]int, len(parts))

	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		numbers[i] = num
	}

	return numbers
}

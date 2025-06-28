package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	print(1)
}

//nolint:gochecknoglobals
var outputStream io.Writer

func initStdStreams() {
	outputStream = os.Stdout
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

func BufferedWrite(w io.Writer, s string) error {
	bw := bufio.NewWriterSize(w, len(s))
	_, err := bw.WriteString(s)
	if err != nil {
		return err
	}
	return bw.Flush()
}

func WriteStringFast(w io.Writer, s string) error {
	buf := unsafe.Slice(unsafe.StringData(s), len(s))
	_, err := w.Write(buf)
	if err != nil {
		return err
	}
	return nil
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

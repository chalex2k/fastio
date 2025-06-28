package main

import (
	"fmt"
	"io"
	"os"
	"strings"
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

//nolint:revive,predeclared
func print(a ...any) {
	var res []string

	for _, v := range a {
		switch val := v.(type) {
		case rune:
			res = append(res, fmt.Sprintf("%c", val))
		default:
			res = append(res, fmt.Sprintf("%v", val))
		}
	}

	_, err := fmt.Fprintln(outputStream, strings.Join(res, " "))
	if err != nil {
		panic(err)
	}
}

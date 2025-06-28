package main

import (
	"fmt"
	"io"
	"os"
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
	for _, v := range a {
		switch val := v.(type) {
		case rune:
			_, err := fmt.Fprintf(outputStream, "%c", val)
			if err != nil {
				panic(err)
			}
		default:
			_, err := fmt.Fprint(outputStream, val)
			if err != nil {
				panic(err)
			}
		}
	}

	_, err := fmt.Fprintln(outputStream) // ln
	if err != nil {
		panic(err)
	}
}

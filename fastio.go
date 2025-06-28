package main

import "fmt"

func main() {
	print(1)
}

//nolint:revive,predeclared
func print(a ...any) {
	fmt.Println(a...)
}

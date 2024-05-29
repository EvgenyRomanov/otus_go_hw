package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	srcString := "Hello, OTUS!"
	fmt.Println(reverse.String(srcString))
}

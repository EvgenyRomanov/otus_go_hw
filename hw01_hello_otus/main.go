package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	srcStrig := "Hello, OTUS!"
	fmt.Println(reverse.String(srcStrig))
}

package main

import (
	"errors"
	"log"
	"os"
)

var ErrCountArgs = errors.New("error count args")

func main() {
	// Place your code here.
	if len(os.Args) < 3 {
		log.Fatal(ErrCountArgs)
	}

	pathToDir := os.Args[1]
	command := os.Args[2:]

	dir, err := ReadDir(pathToDir)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(RunCmd(command, dir))
}

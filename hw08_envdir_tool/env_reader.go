package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	invalidSymbol string = "="
	terminateByte byte   = 0x00
	lineFeed      byte   = '\n'
	trimRightSet         = "\t "
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	// Place your code here
	environments := make(Environment, 10)

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() || !dirEntry.Type().IsRegular() || strings.Contains(dirEntry.Name(), invalidSymbol) {
			continue
		}

		firstLineBytes, err := readLine(dir, dirEntry)
		if err != nil {
			return nil, err
		}
		firstLineString := string(bytes.ReplaceAll(firstLineBytes, []byte{terminateByte}, []byte{lineFeed}))
		firstLineString = strings.TrimRight(firstLineString, trimRightSet)

		envValue := EnvValue{
			Value:      firstLineString,
			NeedRemove: false,
		}

		if firstLineString == "" {
			envValue.NeedRemove = true
		}

		environments[dirEntry.Name()] = envValue
	}

	return environments, nil
}

func readLine(dir string, dirEntry fs.DirEntry) ([]byte, error) {
	f, err := os.OpenFile(filepath.Join(dir, dirEntry.Name()), os.O_RDONLY, 0o644)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	reader := bufio.NewReader(f)
	line, _, err := reader.ReadLine()
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return nil, err
		}
	}

	return line, nil
}

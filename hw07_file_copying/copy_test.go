package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	// Place your code here.
	for _, tCase := range [...]struct {
		in       string
		out      string
		expected string
		offset   int64
		limit    int64
	}{
		{
			in:       "./testdata/input.txt",
			out:      "out_offset0_limit0.txt",
			expected: "./testdata/out_offset0_limit0.txt",
			offset:   0,
			limit:    0,
		},
		{
			in:       "./testdata/input.txt",
			out:      "out_offset0_limit10.txt",
			expected: "./testdata/out_offset0_limit10.txt",
			offset:   0,
			limit:    10,
		},
		{
			in:       "./testdata/input.txt",
			out:      "out_offset0_limit1000.txt",
			expected: "./testdata/out_offset0_limit1000.txt",
			offset:   0,
			limit:    1000,
		},
		{
			in:       "./testdata/input.txt",
			out:      "out_offset0_limit10000.txt",
			expected: "./testdata/out_offset0_limit10000.txt",
			offset:   0,
			limit:    10000,
		},
		{
			in:       "./testdata/input.txt",
			out:      "out_offset100_limit1000.txt",
			expected: "./testdata/out_offset100_limit1000.txt",
			offset:   100,
			limit:    1000,
		},
		{
			in:       "./testdata/input.txt",
			out:      "out_offset6000_limit1000.txt",
			expected: "./testdata/out_offset6000_limit1000.txt",
			offset:   6000,
			limit:    1000,
		},
	} {
		t.Run(fmt.Sprintf("test-%q", tCase.out), func(t *testing.T) {
			f, _ := os.CreateTemp("", tCase.out)
			defer func(name string) {
				_ = os.Remove(name)
			}(f.Name())

			_ = Copy(tCase.in, f.Name(), tCase.offset, tCase.limit)
			out, _ := os.ReadFile(f.Name())
			expected, _ := os.ReadFile(tCase.expected)

			if !bytes.Equal(out, expected) {
				t.Errorf("incoming file and outcomming file not matched")
			}
		})
	}

	t.Run("Unsupported file", func(t *testing.T) {
		err := Copy("/dev/urandom", "output.txt", 0, 0)
		expected := ErrUnsupportedFile
		require.Equal(t, expected, err)
	})

	t.Run("Offset exceeds file size", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "output.txt", 10000, 50)
		expected := ErrOffsetExceedsFileSize
		require.Equal(t, expected, err)
	})
}

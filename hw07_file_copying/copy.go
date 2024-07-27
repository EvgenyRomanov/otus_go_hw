package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.
	fileStat, err := os.Stat(fromPath)
	if err != nil {
		return err
	}

	if !fileStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	fileSize := fileStat.Size()

	if fileSize <= offset {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || limit > (fileSize-offset) {
		limit = fileSize - offset
	}

	fromPathFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	defer func(fromPathFile *os.File) {
		_ = fromPathFile.Close()
	}(fromPathFile)

	toPathFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer func(toPathFile *os.File) {
		_ = toPathFile.Close()
	}(toPathFile)

	_, err = fromPathFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(fromPathFile)

	_, err = io.CopyN(toPathFile, barReader, limit)
	if err != nil {
		return err
	}

	defer bar.Finish()

	return nil
}

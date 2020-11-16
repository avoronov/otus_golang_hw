package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	// ErrUnsupportedFile is raised when usupported file given as source or destination.
	ErrUnsupportedFile = errors.New("unsupported file")
	// ErrOffsetExceedsFileSize is raised when given offset exceeds file size.
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

// Copy copying given as limit amount of bytes from source to destination files starting from offset.
func Copy(fromPath string, toPath string, offset, limit int64) error {
	src, limit, err := prepareSourceAndLimit(fromPath, offset, limit)
	if err != nil {
		return err
	}
	defer src.Close()

	tmpfile, err := ioutil.TempFile("", "temp.dest.*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}()

	err = copy(src, tmpfile, limit)
	if err != nil {
		return err
	}

	err = os.Rename(tmpfile.Name(), toPath)
	if err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	return nil
}

func prepareSourceAndLimit(fromPath string, offset, limit int64) (*os.File, int64, error) {
	src, err := os.Open(fromPath)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to open file: %w", err)
	}

	srcInfo, err := src.Stat()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get file's stat: %w", err)
	}

	if !srcInfo.Mode().IsRegular() {
		return nil, 0, ErrUnsupportedFile
	}

	if offset > srcInfo.Size() {
		return nil, 0, ErrOffsetExceedsFileSize
	}

	remainBytes := srcInfo.Size() - offset
	if limit == 0 || limit > remainBytes {
		limit = remainBytes
	}

	if offset > 0 {
		_, err = src.Seek(offset, 0)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to seek in file: %w", err)
		}
	}

	return src, limit, nil
}

func copy(src io.Reader, dst io.Writer, limit int64) error {
	bar := pb.Start64(limit)
	defer bar.Finish()

	total := int64(0)
	for {
		n, err := io.CopyN(dst, src, 1)
		time.Sleep(time.Millisecond)
		bar.Increment()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("failed to copy between files: %w", err)
		}
		total += n
		if total == limit {
			break
		}
	}

	return nil
}

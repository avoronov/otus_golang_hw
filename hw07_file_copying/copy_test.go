package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	err := Copy("testdata/input.txt", "out.txt", 10000, 100)
	assert.Equal(t, err, ErrOffsetExceedsFileSize)

	err = Copy("/home", "out.txt", 0, 0)
	assert.Equal(t, err, ErrUnsupportedFile)

	Copy("testdata/input.txt", "out.txt", 0, 10)
}

package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

// Environment represents env vars to be set.
type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, errors.Wrapf(err, "got error when read dir %s", dir)
	}

	env := newEnvironment()

	for _, file := range files {
		if !file.Mode().IsRegular() {
			continue
		}

		fileName := file.Name()

		if strings.Contains(fileName, "=") {
			continue
		}

		if file.Size() == 0 {
			env[fileName] = ""
			continue
		}

		val, err := getFirstLine(path.Join(dir, fileName))
		if err != nil {
			return nil, err
		}

		val = sanitizeVal(val)

		env[fileName] = val
	}

	return env, nil
}

func newEnvironment() Environment {
	return make(map[string]string)
}

func getFirstLine(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", errors.Wrapf(err, "got error when read file %s", path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	return scanner.Text(), scanner.Err()
}

func sanitizeVal(s string) string {
	s = strings.TrimRight(s, "\t ")
	s = string(bytes.ReplaceAll([]byte(s), []byte{0x00}, []byte("\n")))

	return s
}

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {

	env, err := ReadDir("./testdata/env")

	require.NoError(t, err)
	expected := Environment(map[string]string{
		"BAR":   "bar",
		"FOO":   "   foo\nwith new line",
		"HELLO": `"hello"`,
		"UNSET": "",
	})
	require.Equal(t, expected, env)

}

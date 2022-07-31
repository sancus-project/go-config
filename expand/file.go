package expand

import (
	"io"
	"os"

	"mvdan.cc/sh/v3/shell"
)

func ExpandString(s string, getEnv func(string) string) (string, error) {
	if getEnv == nil {
		getEnv = os.Getenv
	}

	return shell.Expand(s, getEnv)
}

func ExpandBytes(b []byte, getEnv func(string) string) (string, error) {
	return ExpandString(string(b), getEnv)
}

func ExpandReader(f io.Reader, getEnv func(string) string) (string, error) {
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return ExpandString(string(b[:]), getEnv)
}

func ExpandFile(filename string, getEnv func(string) string) (string, error) {
	var f io.Reader

	if filename == "-" {
		f = os.Stdin
	} else if file, err := os.Open(filename); err != nil {
		return "", err
	} else {
		defer file.Close()
		f = file
	}

	return ExpandReader(f, getEnv)
}

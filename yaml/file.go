package yaml

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"

	"go.sancus.dev/config"
	"go.sancus.dev/config/expand"
)

//
// YAML
//
func LoadString(s string, c interface{}) error {
	s, err := expand.ExpandString(s, nil)
	if err != nil {
		return err
	} else if s == "" {
		// empty file
		return nil
	}

	dec := NewStringDecoder(s)
	dec.KnownFields(true)
	return dec.Decode(c)
}

func LoadBytes(b []byte, c interface{}) error {
	return LoadString(string(b), c)
}

func LoadReader(f io.Reader, c interface{}) error {
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	return LoadString(string(b[:]), c)
}

func LoadFile(filename string, c interface{}) error {
	var f io.Reader

	if filename == "-" {
		f = os.Stdin
	} else if file, err := os.Open(filename); err != nil {
		return err
	} else {
		defer file.Close()
		f = file
	}

	return LoadReader(f, c)
}

func WriteTo(f io.Writer, c interface{}) (int64, error) {
	b, err := yaml.Marshal(c)
	if err != nil {
		// encoding error
		return 0, err
	}

	n, err := f.Write(b)
	return int64(n), err
}

func WriteFile(filename string, c interface{}, filemode os.FileMode) (int64, error) {
	if filemode == 0 {
		filemode = config.DefaultConfigFileMode
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, filemode)
	if err != nil {
		return 0, err
	}

	defer f.Close()
	return WriteTo(f, c)
}

package yaml

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"

	"go.sancus.dev/config"
	"go.sancus.dev/core/errors"
)

//
// YAML
//
func LoadReader(f io.Reader, c interface{}) error {
	if b, err := io.ReadAll(f); err != nil {
		// read error
		return errors.Wrap(err, "ReadAll")
	} else if len(b) == 0 {
		// empty file
		return nil
	} else if err := yaml.Unmarshal(b, c); err != nil {
		// failed to decode
		return errors.Wrap(err, "Unmarshal")
	} else if _, err := config.Validate(c); err != nil {
		// failed to validate
		return errors.Wrap(err, "Validate")
	} else {
		// ready
		return nil
	}
}

func LoadFile(filename string, c interface{}) error {

	if file, err := os.Open(filename); err != nil {
		return err
	} else if err := LoadReader(file, c); err != nil {
		return errors.Wrap(err, "Load: %q", filename)
	} else {
		return nil
	}
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

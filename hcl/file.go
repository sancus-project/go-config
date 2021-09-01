package hcl

import (
	"io"
	"os"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"gopkg.in/dealancer/validate.v2"

	"go.sancus.dev/core/errors"
)

//
// HCL
//
func LoadReader(filename string, f io.Reader, c interface{}) error {
	if b, err := io.ReadAll(f); err != nil {
		// read error
		return errors.Wrap(err, "ReadAll")
	} else if len(b) == 0 {
		// empty file
		return nil
	} else if err := hclsimple.Decode(filename, b, nil, c); err != nil {
		// failed to decode
		return errors.Wrap(err, "Decode")
	} else if err := validate.Validate(c); err != nil {
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
	} else if err := LoadReader(filename, file, c); err != nil {
		return errors.Wrap(err, "Load: %q", filename)
	} else {
		return nil
	}
}

func WriteTo(out io.Writer, c interface{}) (int, error) {
	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(c, f.Body())

	n, err := f.WriteTo(out)
	return int(n), err
}

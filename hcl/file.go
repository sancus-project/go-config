package hcl

import (
	"io"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"gopkg.in/dealancer/validate.v2"

	"go.sancus.dev/core/errors"
)

//
// HCL
//
func LoadReader(filename string, in io.Reader, ctx *hcl.EvalContext, c interface{}) error {
	if b, err := io.ReadAll(in); err != nil {
		// read error
		return errors.Wrap(err, "ReadAll")
	} else if len(b) == 0 {
		// empty file
		return nil
	} else if f, err := hclsyntax.ParseConfig(b, filename, hcl.Pos{Line: 1, Column: 1}); err != nil {
		return errors.Wrap(err, "ParseConfig")
	} else if err := gohcl.DecodeBody(f.Body, ctx, c); err != nil {
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

func LoadFile(filename string, ctx *hcl.EvalContext, c interface{}) error {

	if file, err := os.Open(filename); err != nil {
		return err
	} else if err := LoadReader(filename, file, ctx, c); err != nil {
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

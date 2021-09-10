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
	} else if f, diags := hclsyntax.ParseConfig(b, filename, hcl.Pos{Line: 1, Column: 1}); diags.HasErrors() {
		return errors.New("%s.%s: %w", "hclsyntax", "ParseConfig", diags)
	} else if diags := gohcl.DecodeBody(f.Body, ctx, c); diags.HasErrors() {
		// failed to decode
		return errors.New("%s.%s: %w", "gohcl", "DecodeBody", diags)
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
		return errors.Wrap(err, "%T: %s: %q", c, "LoadFile", filename)
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

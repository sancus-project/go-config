package hcl

import (
	"io"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"gopkg.in/dealancer/validate.v2"

	"go.sancus.dev/config"
	"go.sancus.dev/config/expand"
	"go.sancus.dev/core/errors"
)

//
// HCL
//
func resolveContext(ctx *hcl.EvalContext, key string) string {
	for ctx != nil {
		if v, ok := ctx.Variables[key]; ok {
			return v.AsString()
		}
		ctx = ctx.Parent()
	}
	return os.Getenv(key)
}

func getEnvHandler(ctx *hcl.EvalContext) func(string) string {
	if ctx == nil {
		return nil
	}

	return func(key string) string {
		return resolveContext(ctx, key)
	}
}

func LoadString(filename string, s string, ctx *hcl.EvalContext, c interface{}) (err error) {
	if s, err = expand.ExpandString(s, getEnvHandler(ctx)); err != nil || s == "" {
		return
	} else if f, diags := hclsyntax.ParseConfig([]byte(s), filename, hcl.Pos{Line: 1, Column: 1}); diags.HasErrors() {
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

func LoadBytes(filename string, b []byte, ctx *hcl.EvalContext, c interface{}) (err error) {
	if len(b) == 0 {
		// empty
		return nil
	}

	return LoadString(filename, string(b), ctx, c)
}

func LoadReader(filename string, in io.Reader, ctx *hcl.EvalContext, c interface{}) error {
	if b, err := io.ReadAll(in); err != nil || len(b) == 0 {
		// read error or empty
		return err
	} else {
		return LoadString(filename, string(b[:]), ctx, c)
	}
}

func LoadFile(filename string, ctx *hcl.EvalContext, c interface{}) error {

	if f, err := os.Open(filename); err != nil {
		return err
	} else {
		defer f.Close()

		return LoadReader(filename, f, ctx, c)
	}
}

func WriteTo(out io.Writer, c interface{}) (int64, error) {
	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(c, f.Body())

	return f.WriteTo(out)
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

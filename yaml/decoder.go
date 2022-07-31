package yaml

import (
	"io"
	"strings"

	"gopkg.in/yaml.v3"

	"go.sancus.dev/config"
	"go.sancus.dev/core/errors"
)

type Decoder struct {
	decoder *yaml.Decoder
}

func NewStringDecoder(s string) *Decoder {
	r := strings.NewReader(s)
	return NewDecoder(r)
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		decoder: yaml.NewDecoder(r),
	}
}

func (dec *Decoder) KnownFields(enable bool) {
	dec.decoder.KnownFields(enable)
}

func (dec *Decoder) Decode(v interface{}) (err error) {
	if err := config.SetDefaults(v); err != nil {
		return errors.Wrap(err, "SetDefaults")
	}

	if err := dec.decoder.Decode(v); err != nil {
		return errors.Wrap(err, "Decode")
	}

	if _, err := config.Validate(v); err != nil {
		return errors.Wrap(err, "Validate")
	}

	return nil
}

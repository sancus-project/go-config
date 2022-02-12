package config

import (
	"github.com/creasty/defaults"

	valid "github.com/asaskevich/govalidator"

	"go.sancus.dev/core/errors"
)

func SetDefaults(c interface{}) error {
	return defaults.Set(c)
}

func Validate(c interface{}) (bool, error) {
	return valid.ValidateStruct(c)
}

func Prepare(c interface{}) error {
	if err := SetDefaults(c); err != nil {
		return errors.Wrap(err, "SetDefaults")
	}

	if _, err := Validate(c); err != nil {
		return errors.Wrap(err, "Validate")
	}

	return nil
}

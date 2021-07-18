package flags

import (
	"go.sancus.dev/core/errors"
)

func ErrInvalidVarType(name string, ptr interface{}) error {
	return errors.New("%s: invalid type (%T)", name, ptr)
}

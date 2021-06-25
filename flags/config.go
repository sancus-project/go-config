package flags

import (
	"github.com/creasty/defaults"
)

func SetDefaults(c interface{}) error {
	return defaults.Set(c)
}

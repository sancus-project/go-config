package config

import (
	valid "github.com/asaskevich/govalidator"
)

func Validate(c interface{}) (bool, error) {
	return valid.ValidateStruct(c)
}

//go:generate sh ./types.sh
package flags

import (
	"time"
)

// Uint is a flag of type uint
type Uint interface {
	GetUint() (uint, bool)
}

// GetUint tries to find a field of a given name and return a uint
func (m MapperFunc) GetUint(name string) (uint, bool) {
	var zero uint

	if f := m(name); f != nil {
		if v, ok := f.(Uint); ok {
			return v.GetUint()
		}
	}

	return zero, false
}

// Uint16 is a flag of type uint16
type Uint16 interface {
	GetUint16() (uint16, bool)
}

// GetUint16 tries to find a field of a given name and return a uint16
func (m MapperFunc) GetUint16(name string) (uint16, bool) {
	var zero uint16

	if f := m(name); f != nil {
		if v, ok := f.(Uint16); ok {
			return v.GetUint16()
		}
	}

	return zero, false
}

// Uint32 is a flag of type uint32
type Uint32 interface {
	GetUint32() (uint32, bool)
}

// GetUint32 tries to find a field of a given name and return a uint32
func (m MapperFunc) GetUint32(name string) (uint32, bool) {
	var zero uint32

	if f := m(name); f != nil {
		if v, ok := f.(Uint32); ok {
			return v.GetUint32()
		}
	}

	return zero, false
}

// Int is a flag of type int
type Int interface {
	GetInt() (int, bool)
}

// GetInt tries to find a field of a given name and return a int
func (m MapperFunc) GetInt(name string) (int, bool) {
	var zero int

	if f := m(name); f != nil {
		if v, ok := f.(Int); ok {
			return v.GetInt()
		}
	}

	return zero, false
}

// Int16 is a flag of type int16
type Int16 interface {
	GetInt16() (int16, bool)
}

// GetInt16 tries to find a field of a given name and return a int16
func (m MapperFunc) GetInt16(name string) (int16, bool) {
	var zero int16

	if f := m(name); f != nil {
		if v, ok := f.(Int16); ok {
			return v.GetInt16()
		}
	}

	return zero, false
}

// Int32 is a flag of type int32
type Int32 interface {
	GetInt32() (int32, bool)
}

// GetInt32 tries to find a field of a given name and return a int32
func (m MapperFunc) GetInt32(name string) (int32, bool) {
	var zero int32

	if f := m(name); f != nil {
		if v, ok := f.(Int32); ok {
			return v.GetInt32()
		}
	}

	return zero, false
}

// Bool is a flag of type bool
type Bool interface {
	GetBool() (bool, bool)
}

// GetBool tries to find a field of a given name and return a bool
func (m MapperFunc) GetBool(name string) (bool, bool) {
	var zero bool

	if f := m(name); f != nil {
		if v, ok := f.(Bool); ok {
			return v.GetBool()
		}
	}

	return zero, false
}

// String is a flag of type string
type String interface {
	GetString() (string, bool)
}

// GetString tries to find a field of a given name and return a string
func (m MapperFunc) GetString(name string) (string, bool) {
	var zero string

	if f := m(name); f != nil {
		if v, ok := f.(String); ok {
			return v.GetString()
		}
	}

	return zero, false
}

// StringSlice is a flag of type []string
type StringSlice interface {
	GetStringSlice() ([]string, bool)
}

// GetStringSlice tries to find a field of a given name and return a []string
func (m MapperFunc) GetStringSlice(name string) ([]string, bool) {
	var zero []string

	if f := m(name); f != nil {
		if v, ok := f.(StringSlice); ok {
			return v.GetStringSlice()
		}
	}

	return zero, false
}

// Duration is a flag of type time.Duration
type Duration interface {
	GetDuration() (time.Duration, bool)
}

// GetDuration tries to find a field of a given name and return a time.Duration
func (m MapperFunc) GetDuration(name string) (time.Duration, bool) {
	var zero time.Duration

	if f := m(name); f != nil {
		if v, ok := f.(Duration); ok {
			return v.GetDuration()
		}
	}

	return zero, false
}

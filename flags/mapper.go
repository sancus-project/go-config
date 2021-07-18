package flags

import (
	"time"
)

type Looker interface {
	Lookup(name string) Flag
}

type Mapper interface {
	Looker
	Parse()

	GetUint16(name string) (uint16, bool)
	GetDuration(name string) (time.Duration, bool)

	Var(v interface{}, name string, usage string, args ...interface{}) Mapper
	VarP(v interface{}, name string, short rune, usage string, args ...interface{}) Mapper
}

type Flag interface {
	Changed() bool
}

type mapper func(name string) Flag

type Uint16 interface {
	GetUint16() (uint16, bool)
}

func (m mapper) GetUint16(name string) (uint16, bool) {
	if f := m(name); f != nil {
		if v, ok := f.(Uint16); ok {
			return v.GetUint16()
		}
	}
	return 0, false
}

type Duration interface {
	GetDuration() (time.Duration, bool)
}

func (m mapper) GetDuration(name string) (time.Duration, bool) {
	if f := m(name); f != nil {
		if v, ok := f.(Duration); ok {
			return v.GetDuration()
		}
	}
	return 0, false
}

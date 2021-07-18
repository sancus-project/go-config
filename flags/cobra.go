package flags

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
)

type cobraFlag struct {
	f   *pflag.Flag
	v   interface{}
	out interface{}
}

func (f cobraFlag) Changed() bool {
	return f.f.Changed
}

func (f cobraFlag) Flag() *pflag.Flag {
	return f.f
}

func (f cobraFlag) Raw() interface{} {
	return f.v
}

type CobraMapper struct {
	mapper

	values map[string]*cobraFlag
	set    *pflag.FlagSet
}

func NewCobraMapper(set *pflag.FlagSet) *CobraMapper {
	if set != nil {
		m := &CobraMapper{
			set:    set,
			values: make(map[string]*cobraFlag),
		}
		m.mapper = m.Lookup
		registerMapper(set, m)
		return m
	}
	return nil
}

func (m *CobraMapper) Lookup(name string) Flag {
	if v, ok := m.values[name]; ok {
		return v
	}
	return nil
}

func (m *CobraMapper) addFlag(name string, v interface{}, out interface{}) *CobraMapper {
	p := &cobraFlag{
		f:   m.set.Lookup(name),
		v:   v,
		out: out,
	}
	m.values[name] = p
	return m
}

func (m *CobraMapper) Parse() {
	for _, p := range m.values {
		if p.out == nil {
			// skip
		} else if v, ok := p.GetUint16(); ok {
			// Uint16
			out := p.out.(*uint16)
			*out = v
		} else if v, ok := p.GetDuration(); ok {
			// Duration
			out := p.out.(*time.Duration)
			*out = v
		}
	}
}

// Uint16
func (f cobraFlag) GetUint16() (uint16, bool) {
	var v uint16

	p, ok := f.Raw().(*uint16)
	if ok {
		ok = f.Changed()
		v = *p
	}

	return v, ok
}

func (m *CobraMapper) uint16VarP(out *uint16, v *uint16, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	var zero uint16

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = zero
	m.set.Uint16VarP(v, name, shorthand, zero, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) UintVar16P(out *uint16, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(uint16)
	return m.uint16VarP(out, v, name, shorthand, usage, args...)
}

func (m *CobraMapper) Uint16P(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(uint16)
	return m.uint16VarP(nil, v, name, shorthand, usage, args...)
}

// Bool
func (f cobraFlag) GetBool() (bool, bool) {
	var v bool

	p, ok := f.Raw().(*bool)
	if ok {
		ok = f.Changed()
		v = *p
	}

	return v, ok
}

func (m *CobraMapper) boolVarP(out *bool, v *bool, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	var zero bool

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = zero
	m.set.BoolVarP(v, name, shorthand, zero, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) BoolVarP(out *bool, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(bool)
	return m.boolVarP(out, v, name, shorthand, usage, args...)
}

func (m *CobraMapper) BoolP(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(bool)
	return m.boolVarP(nil, v, name, shorthand, usage, args...)
}

// String
func (f cobraFlag) GetString() (string, bool) {
	var v string

	p, ok := f.Raw().(*string)
	if ok {
		ok = f.Changed()
		v = *p
	}

	return v, ok
}

func (m *CobraMapper) stringVarP(out *string, v *string, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	var zero string

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = zero
	m.set.StringVarP(v, name, shorthand, zero, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) StringVarP(out *string, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(string)
	return m.stringVarP(out, v, name, shorthand, usage, args...)
}

func (m *CobraMapper) StringP(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(string)
	return m.stringVarP(nil, v, name, shorthand, usage, args...)
}

// Duration
func (f cobraFlag) GetDuration() (time.Duration, bool) {
	var v time.Duration

	p, ok := f.Raw().(*time.Duration)
	if ok {
		ok = f.Changed()
		v = *p
	}

	return v, ok
}

func (m *CobraMapper) durationVarP(out *time.Duration, v *time.Duration, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	var zero time.Duration

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}

	*v = zero
	m.set.DurationVarP(v, name, shorthand, zero, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) DurationVarP(out *time.Duration, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(time.Duration)
	return m.durationVarP(out, v, name, shorthand, usage, args...)
}
func (m *CobraMapper) DurationP(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(time.Duration)
	return m.durationVarP(nil, v, name, shorthand, usage, args...)
}

// Magic
func (m *CobraMapper) VarP(ptr interface{}, name string, short rune, usage string, args ...interface{}) Mapper {
	if ptr != nil {
		var shorthand string

		if short != rune(0) {
			shorthand = string(short)
		}

		if len(usage) > 0 && len(args) > 0 {
			usage = fmt.Sprintf(usage, args...)
		}

		if out, ok := ptr.(*uint16); ok {
			v := new(uint16)
			m.uint16VarP(out, v, name, shorthand, usage)
		} else if out, ok := ptr.(*bool); ok {
			v := new(bool)
			m.boolVarP(out, v, name, shorthand, usage)
		} else if out, ok := ptr.(*string); ok {
			v := new(string)
			m.stringVarP(out, v, name, shorthand, usage)
		} else if out, ok := ptr.(*time.Duration); ok {
			v := new(time.Duration)
			m.durationVarP(out, v, name, shorthand, usage)
		} else {
			goto fail
		}
		return m
	}

fail:
	panic(ErrInvalidVarType(name, ptr))
}

func (m *CobraMapper) Var(ptr interface{}, name string, usage string, args ...interface{}) Mapper {
	return m.VarP(ptr, name, 0, usage, args...)
}

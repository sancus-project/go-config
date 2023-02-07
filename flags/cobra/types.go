//go:generate sh ./types.sh
package cobra

import (
	"fmt"
	"time"

	"go.sancus.dev/config/flags"
)

// Parse() updates all fields which mapped flags were used
func (m *CobraMapper) Parse() {
	for _, p := range m.values {
		if p.out == nil {
			// skip
		} else if v, ok := p.GetUint(); ok {
			// Uint
			out := p.out.(*uint)
			*out = v
		} else if v, ok := p.GetUint16(); ok {
			// Uint16
			out := p.out.(*uint16)
			*out = v
		} else if v, ok := p.GetUint32(); ok {
			// Uint32
			out := p.out.(*uint32)
			*out = v
		} else if v, ok := p.GetInt(); ok {
			// Int
			out := p.out.(*int)
			*out = v
		} else if v, ok := p.GetInt16(); ok {
			// Int16
			out := p.out.(*int16)
			*out = v
		} else if v, ok := p.GetInt32(); ok {
			// Int32
			out := p.out.(*int32)
			*out = v
		} else if v, ok := p.GetBool(); ok {
			// Bool
			out := p.out.(*bool)
			*out = v
		} else if v, ok := p.GetString(); ok {
			// String
			out := p.out.(*string)
			*out = v
		} else if v, ok := p.GetDuration(); ok {
			// Duration
			out := p.out.(*time.Duration)
			*out = v
		}
	}
}

// VarP identifies the type of field and calls the right mapper accordingly
func (m *CobraMapper) VarP(ptr interface{}, name string, short rune, usage string, args ...interface{}) flags.Mapper {
	var shorthand string

	if short != rune(0) {
		shorthand = string(short)
	}

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}

	if ptr == nil {
		goto fail
	} else if out, ok := ptr.(*uint); ok {
		v := new(uint)
		m.uintVarP(out, v, name, shorthand, usage)
	} else if out, ok := ptr.(*uint16); ok {
		v := new(uint16)
		m.uint16VarP(out, v, name, shorthand, usage)
	} else if out, ok := ptr.(*uint32); ok {
		v := new(uint32)
		m.uint32VarP(out, v, name, shorthand, usage)
	} else if out, ok := ptr.(*int); ok {
		v := new(int)
		m.intVarP(out, v, name, shorthand, usage)
	} else if out, ok := ptr.(*int16); ok {
		v := new(int16)
		m.int16VarP(out, v, name, shorthand, usage)
	} else if out, ok := ptr.(*int32); ok {
		v := new(int32)
		m.int32VarP(out, v, name, shorthand, usage)
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

fail:
	panic(flags.ErrInvalidVarType(name, ptr))
}

func (m *CobraMapper) Var(ptr interface{}, name string, usage string, args ...interface{}) flags.Mapper {
	return m.VarP(ptr, name, 0, usage, args...)
}

// GetUint() returns a uint if the flag was used
func (f cobraFlag) GetUint() (uint, bool) {
	var v uint

	p, ok := f.Raw().(*uint)
	if ok {
		ok = f.Changed()
		v = *p
	}

	return v, ok
}

func (m *CobraMapper) uintVarP(out *uint, v *uint, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	var zero uint

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = zero
	m.set.UintVarP(v, name, shorthand, zero, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) UintVarP(out *uint, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(uint)
	return m.uintVarP(out, v, name, shorthand, usage, args...)
}

func (m *CobraMapper) UintP(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(uint)
	return m.uintVarP(nil, v, name, shorthand, usage, args...)
}

// GetUint16() returns a uint16 if the flag was used
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

func (m *CobraMapper) Uint16VarP(out *uint16, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(uint16)
	return m.uint16VarP(out, v, name, shorthand, usage, args...)
}

func (m *CobraMapper) Uint16P(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(uint16)
	return m.uint16VarP(nil, v, name, shorthand, usage, args...)
}

// GetUint32() returns a uint32 if the flag was used
func (f cobraFlag) GetUint32() (uint32, bool) {
	var v uint32

	p, ok := f.Raw().(*uint32)
	if ok {
		ok = f.Changed()
		v = *p
	}

	return v, ok
}

func (m *CobraMapper) uint32VarP(out *uint32, v *uint32, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	var zero uint32

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = zero
	m.set.Uint32VarP(v, name, shorthand, zero, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) Uint32VarP(out *uint32, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(uint32)
	return m.uint32VarP(out, v, name, shorthand, usage, args...)
}

func (m *CobraMapper) Uint32P(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(uint32)
	return m.uint32VarP(nil, v, name, shorthand, usage, args...)
}

// GetInt() returns a int if the flag was used
func (f cobraFlag) GetInt() (int, bool) {
	var v int

	p, ok := f.Raw().(*int)
	if ok {
		ok = f.Changed()
		v = *p
	}

	return v, ok
}

func (m *CobraMapper) intVarP(out *int, v *int, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	var zero int

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = zero
	m.set.IntVarP(v, name, shorthand, zero, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) IntVarP(out *int, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(int)
	return m.intVarP(out, v, name, shorthand, usage, args...)
}

func (m *CobraMapper) IntP(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(int)
	return m.intVarP(nil, v, name, shorthand, usage, args...)
}

// GetInt16() returns a int16 if the flag was used
func (f cobraFlag) GetInt16() (int16, bool) {
	var v int16

	p, ok := f.Raw().(*int16)
	if ok {
		ok = f.Changed()
		v = *p
	}

	return v, ok
}

func (m *CobraMapper) int16VarP(out *int16, v *int16, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	var zero int16

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = zero
	m.set.Int16VarP(v, name, shorthand, zero, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) Int16VarP(out *int16, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(int16)
	return m.int16VarP(out, v, name, shorthand, usage, args...)
}

func (m *CobraMapper) Int16P(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(int16)
	return m.int16VarP(nil, v, name, shorthand, usage, args...)
}

// GetInt32() returns a int32 if the flag was used
func (f cobraFlag) GetInt32() (int32, bool) {
	var v int32

	p, ok := f.Raw().(*int32)
	if ok {
		ok = f.Changed()
		v = *p
	}

	return v, ok
}

func (m *CobraMapper) int32VarP(out *int32, v *int32, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	var zero int32

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = zero
	m.set.Int32VarP(v, name, shorthand, zero, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) Int32VarP(out *int32, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(int32)
	return m.int32VarP(out, v, name, shorthand, usage, args...)
}

func (m *CobraMapper) Int32P(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new(int32)
	return m.int32VarP(nil, v, name, shorthand, usage, args...)
}

// GetBool() returns a bool if the flag was used
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

// GetString() returns a string if the flag was used
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

// GetDuration() returns a time.Duration if the flag was used
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

#!/bin/sh
set -eu
if [ -n "${GOFILE:-}" ]; then
	exec > "$GOFILE"
fi

TYPES="
Uint
Uint16
Uint32
Int
Int16
Int32
Bool
String
Duration:time.Duration
"

cat <<EOT
//go:generate sh ./$(basename $0)
package ${GOPACKAGE:-cobra}

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
$(for x in $TYPES; do
		N=${x%:*}
		n=$(echo "$N" | tr 'A-Z' 'a-z')
		if [ "$N" = "$x" ]; then
			T=$(echo "$N" | tr 'A-Z' 'a-z')
		else
			T=${x#*:}
		fi

		cat <<EOT
		} else if v, ok := p.Get$N(); ok {
			// $N
			out := p.out.(*$T)
			*out = v
EOT
done)
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
$(for x in $TYPES; do
		N=${x%:*}
		n=$(echo "$N" | tr 'A-Z' 'a-z')
		if [ "$N" = "$x" ]; then
			T=$n
		else
			T=${x#*:}
		fi

		cat <<EOT
	} else if out, ok := ptr.(*$T); ok {
		v := new($T)
		m.${n}VarP(out, v, name, shorthand, usage)
EOT
done)
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
EOT

for x in $TYPES; do
	N=${x%:*}
	n=$(echo "$N" | tr 'A-Z' 'a-z')
	if [ "$N" = "$x" ]; then
		T=$n
	else
		T=${x#*:}
	fi

cat <<EOT

// Get$N() returns a $T if the flag was used
func (f cobraFlag) Get$N() ($T, bool) {
	var v $T

	p, ok := f.Raw().(*$T)
	if ok {
		ok = f.Changed()
		v = *p
	}

	return v, ok
}

func (m *CobraMapper) ${n}VarP(out *$T, v *$T, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	var zero $T

	if len(usage) > 0 && len(args) > 0 {
		usage = fmt.Sprintf(usage, args...)
	}
	*v = zero
	m.set.${N}VarP(v, name, shorthand, zero, usage)
	return m.addFlag(name, v, out)
}

func (m *CobraMapper) ${N}VarP(out *$T, name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new($T)
	return m.${n}VarP(out, v, name, shorthand, usage, args...)
}

func (m *CobraMapper) ${N}P(name, shorthand string, usage string, args ...interface{}) *CobraMapper {
	v := new($T)
	return m.${n}VarP(nil, v, name, shorthand, usage, args...)
}
EOT
done

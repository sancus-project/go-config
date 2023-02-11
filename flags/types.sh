#!/bin/sh

set -eu
if [ -n "$GOFILE" ]; then
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
StringSlice:[]string
Duration:time.Duration
"

cat <<EOT
//go:generate sh ./$(basename $0)
package ${GOPACKAGE:-flags}

import (
	"time"
)
EOT

for x in $TYPES; do
	N=${x%:*}
	if [ "$N" = "$x" ]; then
		T=$(echo "$N" | tr 'A-Z' 'a-z')
	else
		T=${x#*:}
	fi

	cat <<EOT

// $N is a flag of type $T
type $N interface {
	Get$N() ($T, bool)
}

// Get$N tries to find a field of a given name and return a $T
func (m MapperFunc) Get$N(name string) ($T, bool) {
	var zero $T

	if f := m(name); f != nil {
		if v, ok := f.($N); ok {
			return v.Get$N()
		}
	}

	return zero, false
}
EOT
done

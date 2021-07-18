package cobra

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"go.sancus.dev/config/flags"
)

type (
	Command = cobra.Command
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
	flags.MapperFunc

	values map[string]*cobraFlag
	set    *pflag.FlagSet
}

func NewMapper(set *pflag.FlagSet) *CobraMapper {
	if set != nil {
		m := &CobraMapper{
			set:    set,
			values: make(map[string]*cobraFlag),
		}
		m.MapperFunc = m.Lookup
		flags.RegisterMapper(set, m)
		return m
	}
	return nil
}

func GetMapper(set *pflag.FlagSet) flags.Mapper {
	return flags.GetMapper(set)
}

func (m *CobraMapper) Lookup(name string) flags.Flag {
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

package flags

type Flag interface {
	Changed() bool
}

type MapperFunc func(name string) Flag

type Looker interface {
	Lookup(name string) Flag
}

type Mapper interface {
	Looker
	Parse()

	Var(v interface{}, name string, usage string, args ...interface{}) Mapper
	VarP(v interface{}, name string, short rune, usage string, args ...interface{}) Mapper
}

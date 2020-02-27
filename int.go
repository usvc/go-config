package config

import "github.com/spf13/pflag"

type Int struct {
	Default int
	Value   int
	Base
}

func (s *Int) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     int
		pointer   *int
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().(int); !ok {
		value = *new(int)
	}
	if pointer, ok = s.GetValuePointer().(*int); ok {
		if isZeroValue(shorthand) {
			flags.IntVar(pointer, name, value, usage)
		} else {
			flags.IntVarP(pointer, name, shorthand, value, usage)
		}
	} else {
		if isZeroValue(shorthand) {
			flags.IntP(name, shorthand, value, usage)
		} else {
			flags.Int(name, value, usage)
		}
	}
}

func (u *Int) GetDefault() interface{} {
	return u.Default
}

func (u *Int) GetValuePointer() interface{} {
	return &u.Value
}

func (u *Int) GetValue() interface{} {
	return u.Value
}

func (s *Int) SetValue(value interface{}) {
	s.Value = value.(int)
}

package config

import (
	"github.com/spf13/pflag"
)

type StringSlice struct {
	Shorthand string
	Usage     string
	Default   []string
	Value     []string
	Base
}

func (s *StringSlice) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     []string
		pointer   *[]string
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().([]string); !ok {
		value = *new([]string)
	}
	if pointer, ok = s.GetValuePointer().(*[]string); ok {
		if isZeroValue(shorthand) {
			flags.StringSliceVar(pointer, name, value, usage)
		} else {
			flags.StringSliceVarP(pointer, name, shorthand, value, usage)
		}
	} else {
		if isZeroValue(shorthand) {
			flags.StringSliceP(name, shorthand, value, usage)
		} else {
			flags.StringSlice(name, value, usage)
		}
	}
}

func (u *StringSlice) GetDefault() interface{} {
	return u.Default
}

func (u *StringSlice) GetValuePointer() interface{} {
	return &u.Value
}

func (u *StringSlice) GetValue() interface{} {
	return u.Value
}

func (s *StringSlice) SetValue(value interface{}) {
	s.Value = value.([]string)
}

package config

import "github.com/spf13/pflag"

type IntSlice struct {
	Default []int
	Value   []int
	Base
}

func (s *IntSlice) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     []int
		pointer   *[]int
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().([]int); !ok {
		value = *new([]int)
	}
	if pointer, ok = s.GetValuePointer().(*[]int); ok {
		if isZeroValue(shorthand) {
			flags.IntSliceVar(pointer, name, value, usage)
		} else {
			flags.IntSliceVarP(pointer, name, shorthand, value, usage)
		}
	} else {
		if isZeroValue(shorthand) {
			flags.IntSliceP(name, shorthand, value, usage)
		} else {
			flags.IntSlice(name, value, usage)
		}
	}
}

func (u *IntSlice) GetDefault() interface{} {
	return u.Default
}

func (u *IntSlice) GetValuePointer() interface{} {
	return &u.Value
}

func (u *IntSlice) GetValue() interface{} {
	return u.Value
}

func (s *IntSlice) SetValue(value interface{}) {
	s.Value = value.([]int)
}

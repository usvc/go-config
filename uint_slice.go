package config

import "github.com/spf13/pflag"

type UintSlice struct {
	Shorthand string
	Usage     string
	Default   []uint
	Value     []uint
	Base
}

func (s *UintSlice) ApplyToFlagSet(name string, flags *pflag.FlagSet) {
	var (
		ok        bool
		value     []uint
		pointer   *[]uint
		shorthand = s.GetShorthand()
		usage     = s.GetUsage()
	)
	if value, ok = s.GetDefault().([]uint); !ok {
		value = *new([]uint)
	}
	if pointer, ok = s.GetValuePointer().(*[]uint); ok {
		if isZeroValue(shorthand) {
			flags.UintSliceVar(pointer, name, value, usage)
		} else {
			flags.UintSliceVarP(pointer, name, shorthand, value, usage)
		}
	} else {
		if isZeroValue(shorthand) {
			flags.UintSliceP(name, shorthand, value, usage)
		} else {
			flags.UintSlice(name, value, usage)
		}
	}
}

func (u *UintSlice) GetDefault() interface{} {
	return u.Default
}

func (u *UintSlice) GetValuePointer() interface{} {
	return &u.Value
}

func (u *UintSlice) GetValue() interface{} {
	return u.Value
}

func (s *UintSlice) SetValue(value interface{}) {
	s.Value = value.([]uint)
}

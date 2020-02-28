package config

import "github.com/spf13/pflag"

type UintSlice struct {
	Shorthand string
	Usage     string
	Default   []uint
	Value     []uint
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
	pointer = s.GetValuePointer().(*[]uint)
	if isZeroValue(shorthand) {
		flags.UintSliceVar(pointer, name, value, usage)
	} else {
		flags.UintSliceVarP(pointer, name, shorthand, value, usage)
	}
}

func (s *UintSlice) GetDefault() interface{} {
	return s.Default
}

func (s *UintSlice) GetShorthand() string {
	return s.Shorthand
}

func (s *UintSlice) GetUsage() string {
	return s.Usage
}

func (s *UintSlice) GetValuePointer() interface{} {
	return &s.Value
}

func (s *UintSlice) GetValue() interface{} {
	return s.Value
}

func (s *UintSlice) SetValue(value interface{}) {
	s.Value = value.([]uint)
}
